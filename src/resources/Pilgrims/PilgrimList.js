import * as React from "react";
import {
    List,
    Datagrid,
    TextField,
    EmailField,
    NumberField,
} from 'react-admin';

export const PilgrimList = props => (
    <List {...props} >
        <Datagrid>
            <TextField source="id" label="PilgrimID" />
            <TextField source="username" label="Username" />
            <EmailField source="email" label="Email" />
        </Datagrid>
    </List>
);