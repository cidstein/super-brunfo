DOCKER_REGISTRY:=006869236742.dkr.ecr.us-east-2.amazonaws.com

UUID:=$(shell uuidgen)

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o _cicd/dist/super-brunfo github.com/cidstein/super-brunfo

.PHONY: docker/build
docker/build: build
	docker build -t super-brunfo:latest -t $(DOCKER_REGISTRY)/super-brunfo:$(UUID) _cicd

.PHONY: docker/push
docker/push: docker/build
	docker push $(DOCKER_REGISTRY)/super-brunfo:$(UUID)

.PHONY: release
deploy: docker/push
	aws ecs register-task-definition \
		--family super-brunfo \
		--requires-compatibilities "FARGATE" \
		--container-definitions "[{\
			\"name\": \"app\", \
			\"image\": \"$(DOCKER_REGISTRY)/super-brunfo:$(UUID)\", \
			\"essential\": true, \
			\"command\": [\"/app/bin/super-brunfo\"], \
			\"portMappings\": [{ \
				\"containerPort\": 8080, \
				\"hostPort\": 8080, \
				\"protocol\": \"tcp\" \
			}], \
			\"logConfiguration\": { \
				\"logDriver\": \"awslogs\", \
				\"options\": { \
					\"awslogs-create-group\": \"true\", \
					\"awslogs-group\": \"/ecs/super-brunfo\", \
					\"awslogs-region\": \"us-east-2\", \
					\"awslogs-stream-prefix\": \"ecs\" \
				}, \
				\"secretOptions\": [] \
			}, \
			\"secrets\": [ \
				{ \
					\"name\": \"SUPER_BRUNFO_SECRET\", \
					\"valueFrom\": \"arn:aws:secretsmanager:us-east-2:006869236742:secret:Testing/SuperBrunfo-mfD2VF\" \
				} \
			] \
		}]" \
		--cpu 512 \
		--memory 2048 \
		--runtime-platform "{ \
			\"cpuArchitecture\": \"X86_64\", \
			\"operatingSystemFamily\": \"LINUX\" \
		}" \
		--execution-role-arn "arn:aws:iam::006869236742:role/ecsTaskExecutionRole" \
		--network-mode "awsvpc"