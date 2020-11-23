import * as React from "react";
import { Route } from 'react-router-dom';
import RegisterPage from './components/Register/RegisterPage';

export default [
    <Route exact path="/register" component={RegisterPage} noLayout />,
];