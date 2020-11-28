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
} from 'react-admin';

export const TicketList = props => (
    <List {...props} >
        <Datagrid>
            <NumberField source="id" label="TicketID" />
            <EmailField source="email" label="Email" />
            <TextField source="topic" label="Topic" />
            <TextField source="details" label="Details" />
            <TextField source="status" label="Status" />
        </Datagrid>
    </List>
);

export const TicketCreate = props => (
    <Create {...props}>
        <SimpleForm>
            <TextInput source="email" label="Email" />
            <TextInput source="topic" label="Topic" />
            <TextInput source="details" fullWidth='true' label="Details" />
        </SimpleForm>
    </Create>
);