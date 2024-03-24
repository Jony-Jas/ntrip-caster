/* eslint-disable react/prop-types */
import active from "./assets/active.png";

function Details({bs,users}) {
  return (
    <div style={{ display: "flex", align: "center" }}>
      <div style={{ width: "50%" }}>
        <h2 style={{ display: "inline-block" }}>Active BaseStations</h2>
        <img src={active} alt="active" className="status-icon" />
        <div>
          {bs?.map((bs) => (
            <li key={bs.Name}>
              {bs.Name} - ({bs.PosX},{bs.PosY})
            </li>
          ))}
        </div>
      </div>
      <div style={{ width: "50%" }}>
        <h2 style={{ display: "inline-block" }}>Active Users</h2>
        <img src={active} alt="active" className="status-icon" />
        <div>
          {users?.map((u) => (
            <li key={u.Name}>
              {u.Name} - ({u.PosX},{u.PosY})
            </li>
          ))}
        </div>
      </div>
    </div>
  );
}

export default Details;
