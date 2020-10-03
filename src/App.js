import * as React from "react";
import './App.css';
import history from "./history";
import customRoutes from './customRoutes';
import { Admin, Resource, ListGuesser } from 'react-admin';

import LoginPage from './components/Login/LoginPage';
import LogoutButton from './components/Logout/LogoutButton';

import DataProvider from './api/DataProvider/DataProvider';
import AuthProvider from './api/AuthProvider/AuthProvider';

import LibraryBooksIcon from '@material-ui/icons/LibraryBooks';
import ContactSupportIcon from '@material-ui/icons/ContactSupport';

function App() {
  return (
    <Admin customRoutes={customRoutes} loginPage={LoginPage} logoutButton={LogoutButton} history={history} authProvider={AuthProvider} dataProvider={DataProvider}>
      <Resource name="projects" icon={LibraryBooksIcon} list={ListGuesser} />
      <Resource name="tickets" icon={ContactSupportIcon} list={ListGuesser} />
    </Admin>
  );
}

export default App;
