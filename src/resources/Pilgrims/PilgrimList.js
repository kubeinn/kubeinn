import * as React from "react";
import {
    List,
    Create,
    Datagrid,
    TextField,
    EmailField,
    NumberField,
    SimpleForm,
    TextInput,
    PasswordInput,
} from 'react-admin';

export const PilgrimList = props => (
    <List {...props} >
        <Datagrid>
            <TextField source="id" label="PilgrimID" />
            <TextField source="username" label="Username" />
            <EmailField source="email" label="Email" />
            <TextField source="regcode" label="Registration Code" />
        </Datagrid>
    </List>
);

export const PilgrimCreate = props => (
    <Create {...props}>
        <SimpleForm>
            <TextInput source="username" label="Username" />
            <TextInput source="email" label="Email" />
        </SimpleForm>
    </Create>
);