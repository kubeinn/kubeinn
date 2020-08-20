import * as React from "react";
import './App.css';
import { Admin, Resource } from 'react-admin';
import simpleRestProvider from 'ra-data-simple-rest';
// import customRoutes from './routes';

import Dashboard from './Components/Dashboard/Dashboard';
import AuthProvider from './Components/AuthProvider/AuthProvider';
import { UserList } from './Resources/Users/UserList';
import { PostList, PostEdit } from './Resources/Posts/PostList';
import LibraryBooksIcon from '@material-ui/icons/LibraryBooks';
import ContactSupportIcon from '@material-ui/icons/ContactSupport';

const dataProvider = simpleRestProvider('http://localhost/api/pilgrim');

function App() {
  return (
    <Admin dashboard={Dashboard} authProvider={AuthProvider} dataProvider={dataProvider}>
      <Resource name="users" list={UserList} />
      <Resource name="posts" list={PostList} edit={PostEdit} />
      {/* <Resource name="projects" icon={LibraryBooksIcon} list={UserList} />
      <Resource name="tickets" icon={ContactSupportIcon} list={UserList} /> */}
    </Admin>
  );
}

export default App;
