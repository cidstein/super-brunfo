import { Form, Button, Image, Container, Col, Row } from 'react-bootstrap';

export default function SignUp() {
  function handleSubmit(event: any) {
    event.preventDefault();
    const data = new FormData(event.target);

    fetch('http://localhost:3000/api/signup', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': '*',
      },
      body: data,
    });
  }

  return (
    <Container className="container">
      <Row className="show-grid">
        <Col md={{ span: 4, offset: 4 }} className="border-center">
          <Image src="/logo.png" alt="logo" className='img-logo' />
        </Col>
      </Row>
      <Row className="show-grid">
        <Col md={{ span: 4, offset: 4 }} className="border-center">
          <Form>
            <Form.Group className="mb-3" controlId="formBasicEmail">
              <Form.Label>Email address</Form.Label>
              <Form.Control type="email" placeholder="Enter email" />
            </Form.Group>

            <Form.Group className="mb-3" controlId="formBasicPassword">
              <Form.Label>Password</Form.Label>
              <Form.Control type="password" placeholder="Password" />
            </Form.Group>

            <Form.Group className="mb-3" controlId="formBasicPassword">
              <Form.Label>Nickname</Form.Label>
              <Form.Control type="nickname" placeholder="Nickname" />
            </Form.Group>
            <Button variant="dark" type="submit">
              Register
            </Button>
          </Form>
        </Col>
      </Row>
    </Container>
  );
}
