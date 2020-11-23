import * as React from "react";
import './App.css';
import history from "./history";
import { Admin, Resource } from 'react-admin';
import customRoutes from './customRoutes';

import DataProvider from './api/DataProvider/DataProvider';
import AuthProvider from './api/AuthProvider/AuthProvider';
import { InnkeeperList } from './resources/Innkeepers/Innkeeper';
import { PilgrimList } from './resources/Pilgrims/PilgrimList';
import { VillageCreate, VillageList } from './resources/Villages/VillageList';
import { TicketList, TicketCreate, TicketEdit } from './resources/Tickets/TicketList';
import { ReeveList, ReeveEdit } from './resources/Reeves/ReeveList';
import { ProjectList } from './resources/Projects/ProjectList';

import LoginPage from './components/Login/LoginPage';
import LogoutButton from './components/Logout/LogoutButton';

import HouseIcon from '@material-ui/icons/House';
import FaceIcon from '@material-ui/icons/Face';
import SupervisorAccountIcon from '@material-ui/icons/SupervisorAccount';
import SupervisedUserCircleIcon from '@material-ui/icons/SupervisedUserCircle';
import LibraryBooksIcon from '@material-ui/icons/LibraryBooks';
import ContactSupportIcon from '@material-ui/icons/ContactSupport';

function App() {
  return (
    <Admin customRoutes={customRoutes} loginPage={LoginPage} logoutButton={LogoutButton} history={history} authProvider={AuthProvider} dataProvider={DataProvider}>
      <Resource name="innkeepers" icon={FaceIcon} list={InnkeeperList} />
      <Resource name="villages" icon={HouseIcon} list={VillageList} create={VillageCreate} />
      <Resource name="reeves" icon={SupervisorAccountIcon} list={ReeveList} edit={ReeveEdit} />
      <Resource name="pilgrims" icon={SupervisedUserCircleIcon} list={PilgrimList} />
      <Resource name="projects" icon={LibraryBooksIcon} list={ProjectList} />
      <Resource name="tickets" icon={ContactSupportIcon} list={TicketList} create={TicketCreate} edit={TicketEdit} />
    </Admin>
  );
}

export default App;
