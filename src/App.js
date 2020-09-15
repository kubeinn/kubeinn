import * as React from "react";
import './App.css';
import { Admin, Resource, ListGuesser } from 'react-admin';
import simpleRestProvider from 'ra-data-simple-rest';
import createHistory from 'history/createBrowserHistory';
// import customRoutes from './routes';

import Dashboard from './Components/Dashboard/Dashboard';
import AuthProvider from './Components/AuthProvider/AuthProvider';

import GroupIcon from '@material-ui/icons/Group';
import SupervisorAccountIcon from '@material-ui/icons/SupervisorAccount';
import SupervisedUserCircleIcon from '@material-ui/icons/SupervisedUserCircle';
import LibraryBooksIcon from '@material-ui/icons/LibraryBooks';
import ContactSupportIcon from '@material-ui/icons/ContactSupport';

const dataProvider = simpleRestProvider(window._env_.KUBEINN_POSTGREST_URL);
const history = createHistory({ basename: 'innkeeper' });

function App() {
  return (
    <Admin history={history} dashboard={Dashboard} authProvider={AuthProvider} dataProvider={dataProvider}>
      <Resource name="villages" icon={GroupIcon} list={ListGuesser} />
      <Resource name="pilgrims" icon={SupervisedUserCircleIcon} list={ListGuesser} />
      <Resource name="innkeepers" icon={SupervisorAccountIcon} list={ListGuesser} />
      <Resource name="projects" icon={LibraryBooksIcon} list={ListGuesser} />
      <Resource name="tickets" icon={ContactSupportIcon} list={ListGuesser} />
    </Admin>
  );
}

export default App;
