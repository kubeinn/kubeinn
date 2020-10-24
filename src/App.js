import * as React from "react";
import './App.css';
import history from "./history";
import { Admin, Resource } from 'react-admin';
import customRoutes from './customRoutes';

import DataProvider from './api/DataProvider/DataProvider';
import AuthProvider from './api/AuthProvider/AuthProvider';
import { TicketList, TicketCreate } from './resources/Tickets/TicketList';
import { ProjectList, ProjectCreate } from './resources/Projects/ProjectList';

import LoginPage from './components/Login/LoginPage';
import LogoutButton from './components/Logout/LogoutButton';

import LibraryBooksIcon from '@material-ui/icons/LibraryBooks';
import ContactSupportIcon from '@material-ui/icons/ContactSupport';

function App() {
  return (
    <Admin customRoutes={customRoutes} loginPage={LoginPage} logoutButton={LogoutButton} history={history} authProvider={AuthProvider} dataProvider={DataProvider}>
      <Resource name="projects" icon={LibraryBooksIcon} list={ProjectList} create={ProjectCreate} />
      <Resource name="tickets" icon={ContactSupportIcon} list={TicketList} create={TicketCreate} />
    </Admin>
  );
}

export default App;
