import * as React from "react";
import {
    List,
    Datagrid,
    TextField,
    NumberField,
} from 'react-admin';

export const ProjectList = props => (
    <List {...props} >
        <Datagrid>
            <NumberField source="id" label="ProjectID" />
            <TextField source="pilgrimid" label="PilgrimID" />
            <TextField source="title" label="Title" />
            <TextField source="details" label="Details" />
            <NumberField source="cpu" label="CPU Limits (number)" />
            <NumberField source="memory" label="Memory Limits (bytes)" />
            <NumberField source="storage" label="Storage Requests (bytes)" />
        </Datagrid>
    </List>
);