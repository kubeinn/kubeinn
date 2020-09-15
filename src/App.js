import * as React from "react";
import './App.css';
import { Admin, Resource, ListGuesser } from 'react-admin';
import simpleRestProvider from 'ra-data-simple-rest';
import createHistory from 'history/createBrowserHistory';
// import customRoutes from './routes';

import AuthProvider from './Components/AuthProvider/AuthProvider';
// import { UserList } from './Resources/Users/UserList';
// import { PostList, PostEdit } from './Resources/Posts/PostList';


import LibraryBooksIcon from '@material-ui/icons/LibraryBooks';
import ContactSupportIcon from '@material-ui/icons/ContactSupport';

const dataProvider = simpleRestProvider(window._env_.KUBEINN_POSTGREST_URL);
const history = createHistory({ basename: 'pilgrim' });

function App() {
  return (
    <Admin history={history} authProvider={AuthProvider} dataProvider={dataProvider}>
      <Resource name="projects" icon={LibraryBooksIcon} list={ListGuesser} />
      <Resource name="tickets" icon={ContactSupportIcon} list={ListGuesser} />
    </Admin>
  );
}

export default App;
