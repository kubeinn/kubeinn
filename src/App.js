import * as React from "react";
import './App.css';
import history from "./history";
import { Admin, Resource } from 'react-admin';
// import customRoutes from './routes';

import Dashboard from './resources/Dashboard/Dashboard';
import DataProvider from './api/DataProvider/DataProvider';
import AuthProvider from './api/AuthProvider/AuthProvider';
import { InnkeeperList, InnkeeperEdit, InnkeeperCreate } from './resources/Innkeepers/Innkeeper';
import { PilgrimList } from './resources/Pilgrims/PilgrimList';
import { VillageList } from './resources/Villages/VillageList';
import { TicketList } from './resources/Tickets/TicketList';
import { ProjectList } from './resources/Projects/ProjectList';

import HouseIcon from '@material-ui/icons/House';
import SupervisorAccountIcon from '@material-ui/icons/SupervisorAccount';
import SupervisedUserCircleIcon from '@material-ui/icons/SupervisedUserCircle';
import LibraryBooksIcon from '@material-ui/icons/LibraryBooks';
import ContactSupportIcon from '@material-ui/icons/ContactSupport';

const dataProvider = DataProvider;

function App() {
  return (
    <Admin history={history} dashboard={Dashboard} authProvider={AuthProvider} dataProvider={dataProvider}>
      <Resource name="villages" icon={HouseIcon} list={VillageList} />
      <Resource name="pilgrims" icon={SupervisedUserCircleIcon} list={PilgrimList} />
      <Resource name="innkeepers" icon={SupervisorAccountIcon} list={InnkeeperList} create={InnkeeperCreate} />
      <Resource name="projects" icon={LibraryBooksIcon} list={ProjectList} />
      <Resource name="tickets" icon={ContactSupportIcon} list={TicketList} />
    </Admin>
  );
}

export default App;
