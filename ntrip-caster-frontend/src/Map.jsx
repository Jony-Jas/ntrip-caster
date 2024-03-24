import {
  MapContainer,
  TileLayer,
  Marker,
  Popup,
  Polyline,
} from "react-leaflet";
import "leaflet/dist/leaflet.css";
import L from "leaflet";
import userMarkerImage from "./assets/user-placeholder.png";
import bsMarkerImage from "./assets/station.png";

const bsIcon = new L.Icon({
  iconUrl: bsMarkerImage,
  iconSize: [40, 40], // Set the size of your custom marker image
  iconAnchor: [16, 32], // Adjust the anchor point if needed
  popupAnchor: [0, -322], // Adjust the popup anchor if needed
});

const userIcon = new L.Icon({
  iconUrl: userMarkerImage,
  iconSize: [32, 32],
  iconAnchor: [16, 32],
  popupAnchor: [0, -32],
});

// eslint-disable-next-line react/prop-types
function Map({ bs, users }) {
  console.log(bs);
  console.log(users);

  const generateBsMarker = () => {
    if(!bs) return;
    // eslint-disable-next-line react/prop-types
    const bsMarkers = bs.map((b) => (
      <Marker position={[b.PosX, b.PosY]} key={b.Name} icon={bsIcon}>
        <Popup>{b.Name}</Popup>
      </Marker>
    ));
    return bsMarkers;
  };

  const generateUserMarker = () => {
    if(!users) return;
    // eslint-disable-next-line react/prop-types
    const userMarkers = users.map((user) => (
      <Marker position={[user.PosX, user.PosY]} key={user.Name} icon={userIcon}>
        <Popup>{user.Name}</Popup>
      </Marker>
    ));
    return userMarkers;
  };

  const drawPolyline = () => {
    if (!bs || !users) return;
    // eslint-disable-next-line react/prop-types
    const coordinates = users.map((user) => {
      if (user.Bs === "")
        return [
          [user.PosX, user.PosY],
          [user.PosX, user.PosY],
        ];

      // eslint-disable-next-line react/prop-types
      const bsPosition = bs.find((b) => b.Name === user.Bs);
      const line = [
        [user.PosX, user.PosY],
        [bsPosition.PosX, bsPosition.PosY],
      ];
      return line;
    });

    console.log(coordinates);

    const polylines = coordinates.map((c) => (
      <Polyline positions={c} color="blue" dashArray={[10, 10]} key={c} />
    ));

    return polylines;
  };

  return (
    <div>
      <MapContainer
        center={[3, 3]}
        zoom={5}
        style={{ height: "400px", width: "100%" }}
      >
        {/* TileLayer for the base map */}
        <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />

        {generateBsMarker()}
        {generateUserMarker()}
        {drawPolyline()}
        {/* <Polyline positions={coordinates} color="blue" dashArray={[10, 10]} /> */}
      </MapContainer>
    </div>
  );
}

export default Map;
