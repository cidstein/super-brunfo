import * as React from 'react';

// import { Button, ButtonGroup, OverlayTrigger, Stack, Tooltip } from 'react-bootstrap';

// import Card from 'react-bootstrap/Card';
// import ListGroup from 'react-bootstrap/ListGroup';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faRocket, faBrain, faShield, faDumbbell, faRunning } from '@fortawesome/free-solid-svg-icons'

import { Card } from 'primereact/card';
import { Button } from 'primereact/button';
import { Divider } from 'primereact/divider';
import { Image } from 'primereact/image';

interface CardProps {
  name: string;
  attack: string;
  defense: string;
  intelligence: string;
  agility: string;
  resilience: string;
  imageURL: string;
}

export default function ActionAreaCard(props: CardProps) {
  const { name, attack, defense, intelligence, agility, resilience, imageURL } = props;

  const header = <h1>{name}</h1>;
  const footer = <span>
    Lorem Ipsum is simply dummy text of the printing and typesetting industry.
  </span>;

  return (
    <Card
      header={header}
      footer={footer}
      style={{
        width: '20rem',
        backgroundColor: 'rgba(255, 255, 255, 0.5)',
      }}
    >
      <Image alt="Card" src={imageURL} imageClassName="customCardImg" />

      <div className="flex">
        <FontAwesomeIcon icon={faRocket} />
        <FontAwesomeIcon icon={faShield} />
        <FontAwesomeIcon icon={faBrain} />
        <FontAwesomeIcon icon={faRunning} />
        <FontAwesomeIcon icon={faDumbbell} />
      </div>
    </Card >


    // <Card bg='light' style={{ width: '17rem' }}>

    //   <Card.Img
    //     src="data:image/jpeg;base64,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"
    //     bsPrefix='customCardImgOverlay'
    //   />
    //   <Card.ImgOverlay>
    //     <Card.Header as="h5">{name}</Card.Header>

    //     {/* <Card.Body> */}
    //     <Card.Img src={imageURL} bsPrefix="customCardImg" />

    //     <Stack
    //       direction="horizontal"
    //       bsPrefix="customStack"
    //     >
    //       <Card.Body>
    //         <OverlayTrigger
    //           key={name}
    //           placement='top'
    //           overlay={
    //             <Tooltip id={`tooltip-${attack}`}>
    //               Attack <strong>{attack}</strong>
    //             </Tooltip>
    //           }
    //         >
    //           <FontAwesomeIcon icon={faRocket} />
    //         </OverlayTrigger>
    //         <Card.Text>{attack}</Card.Text>
    //       </Card.Body>
    //       <Card.Body>
    //         <OverlayTrigger
    //           key={name}
    //           placement='top'
    //           overlay={
    //             <Tooltip id={`tooltip-${defense}`}>
    //               Defense <strong>{defense}</strong>
    //             </Tooltip>
    //           }
    //         >
    //           <FontAwesomeIcon icon={faShield} />
    //         </OverlayTrigger>
    //         <Card.Text>{defense}</Card.Text>
    //       </Card.Body>
    //       <Card.Body>
    //         <OverlayTrigger
    //           key={name}
    //           placement='top'
    //           overlay={
    //             <Tooltip id={`tooltip-${intelligence}`}>
    //               Intelligence <strong>{intelligence}</strong>
    //             </Tooltip>
    //           }
    //         >
    //           <FontAwesomeIcon icon={faBrain} />
    //         </OverlayTrigger>
    //         <Card.Text>{intelligence}</Card.Text>
    //       </Card.Body>
    //       <Card.Body>
    //         <OverlayTrigger
    //           key={name}
    //           placement='top'
    //           overlay={
    //             <Tooltip id={`tooltip-${agility}`}>
    //               Agility <strong>{agility}</strong>
    //             </Tooltip>
    //           }
    //         >
    //           <FontAwesomeIcon icon={faRunning} />
    //         </OverlayTrigger>
    //         <Card.Text>{agility}</Card.Text>
    //       </Card.Body>
    //       <Card.Body>
    //         <OverlayTrigger
    //           key={name}
    //           placement='top'
    //           overlay={
    //             <Tooltip id={`tooltip-${resilience}`}>
    //               Resilience <strong>{resilience}</strong>
    //             </Tooltip>
    //           }
    //         >
    //           <FontAwesomeIcon icon={faDumbbell} />
    //         </OverlayTrigger>
    //         <Card.Text>{resilience}</Card.Text>
    //       </Card.Body>
    //     </Stack>

    //     <Card.Text style={{
    //       paddingTop: '0px',
    //       marginTop: '0px',
    //       marginLeft: '8px',
    //       marginRight: '8px',
    //       fontSize: '11px',
    //       display: 'grid',
    //       placeItems: 'center',
    //     }}>
    //       Lorem Ipsum is simply dummy text of the printing and typesetting industry.
    //     </Card.Text>

    //     {/* </Card.Body> */}
    //   </Card.ImgOverlay>
    // </Card>
  );
}
