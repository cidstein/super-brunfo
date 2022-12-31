interface CardProps {
  name: string;
  attack: string;
  defense: string;
  intelligence: string;
  agility: string;
  resilience: string;
  flavourText: string;
  imageURL: string;
}

export default function ActionAreaCard(props: CardProps) {
  const { name, attack, defense, intelligence, agility, resilience, flavourText, imageURL } = props;

  return (
    <div className="card-container">
      <div className="card-background">

        <div className="card-frame">

          <div className="frame-header">
            <h1 className="name">{name}</h1>
            <i className="ms ms-g" id="mana-icon"></i>
          </div>

          <img className="frame-art" src={imageURL} alt={name} />

          <div className="frame-type-line">
            <h1 className="type">Legendary Enchantment</h1>
            {/* <img src="https://image.ibb.co/kzaLjn/OGW_R.png" id="set-icon" alt="OGW-icon"> */}
          </div>

          <div className="frame-text-box">
            {/* <p className="description ftb-inner-margin">Cosmic Dragon is a legendary creature with flying and haste abilities
            </p> */}
            <p className="flavour-text">{flavourText}</p>
          </div>

          <div className="frame-bottom-info inner-margin">
            <div className="fbi-left">
              <p>140/184 R</p>
              {/* <p>OGW &#x2022; EN <!--   <img className="paintbrush" src="https://image.ibb.co/e2VxAS/paintbrush_white.png" alt="paintbrush icon">--> Wesley Burt</p> */}
            </div>

            <div className="fbi-center"></div>

            <div className="fbi-right">
              &#169; 2022 Creatures of infinity
            </div>
          </div>
        </div>
      </div>
    </div >
  );
}
