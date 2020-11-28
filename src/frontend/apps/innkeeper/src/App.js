import * as React from "react";
import './App.css';
import history from "./history";
import { Admin, Resource } from 'react-admin';
import customRoutes from './customRoutes';

import DataProvider from './api/DataProvider/DataProvider';
import AuthProvider from './api/AuthProvider/AuthProvider';
import { InnkeeperCreate, InnkeeperEdit, InnkeeperList } from './resources/Innkeepers/Innkeeper';
import { PilgrimList, PilgrimCreate, PilgrimEdit } from './resources/Pilgrims/PilgrimList';
import { TicketList, TicketCreate, TicketEdit } from './resources/Tickets/TicketList';
import { ProjectList, ProjectCreate } from './resources/Projects/ProjectList';

import LoginPage from './components/Login/LoginPage';
import LogoutButton from './components/Logout/LogoutButton';

import FaceIcon from '@material-ui/icons/Face';
import SupervisedUserCircleIcon from '@material-ui/icons/SupervisedUserCircle';
import LibraryBooksIcon from '@material-ui/icons/LibraryBooks';
import ContactSupportIcon from '@material-ui/icons/ContactSupport';

function App() {
  return (
    <Admin customRoutes={customRoutes} loginPage={LoginPage} logoutButton={LogoutButton} history={history} authProvider={AuthProvider} dataProvider={DataProvider}>
      <Resource name="innkeepers" icon={FaceIcon} list={InnkeeperList} create={InnkeeperCreate} edit={InnkeeperEdit} />
      <Resource name="pilgrims" icon={SupervisedUserCircleIcon} list={PilgrimList} create={PilgrimCreate} edit={PilgrimEdit} />
      <Resource name="projects" icon={LibraryBooksIcon} list={ProjectList} create={ProjectCreate} />
      <Resource name="tickets" icon={ContactSupportIcon} list={TicketList} create={TicketCreate} edit={TicketEdit} />
    </Admin>
  );
}

export default App;
