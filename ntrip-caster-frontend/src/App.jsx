import { useEffect, useState } from "react";
import Map from "./Map";
import useWebSocket from "react-use-websocket";
import Header from "./Header";
import Details from "./Details";

const App = () => {
  const socketUrl = "ws://localhost:8080/admin";
  const { lastMessage, readyState } = useWebSocket(socketUrl);

  const[baseStations,setBaseStations] = useState([]);
  const [users,setUsers] = useState([]);

  useEffect(()=>{
    if(lastMessage){
      const data = JSON.parse(lastMessage.data);
      setBaseStations(data.BaseStations);
      setUsers(data.Users);
    }
  },[lastMessage])

  return (
    <div className="container">
      <Header readyState={readyState}/>
      <Map bs={baseStations} users={users} />
      <Details bs={baseStations} users={users} />
    </div>
  );
};

export default App;
