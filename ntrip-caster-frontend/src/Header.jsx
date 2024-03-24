import active from "./assets/active.png";
import inactive from "./assets/inactive.png";

// eslint-disable-next-line react/prop-types
function Header({ readyState }) {
  return (
    <div>
      <h1 className="header">NTRIP CASTER</h1>
      <span>
        {readyState === 1 ? (
          <>
            <img src={active} alt="active" className="status-icon" />
            <span>Connected</span>
          </>
        ) : (
          <>
            <img src={inactive} alt="active" className="status-icon" />
            <span>Disconnected</span>
          </>
        )}
      </span>
    </div>
  );
}

export default Header;
