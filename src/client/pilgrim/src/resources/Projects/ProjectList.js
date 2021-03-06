import * as React from "react";
import {
    List,
    Create,
    Datagrid,
    TextField,
    NumberField,
    SimpleForm,
    TextInput,
    NumberInput,
    Toolbar,
    SaveButton,
    DeleteButton
} from 'react-admin';
import { Button } from '@material-ui/core';

function copyToClipboard(text) {
    const dummy = document.createElement('textarea');
    dummy.innerHTML = text;
    const parentElement = document.getElementById('root');
    parentElement.appendChild(dummy);
    dummy.select();
    document.execCommand('copy');
    parentElement.removeChild(dummy);
}

const CopyKubecfgField = ({ source, record = {} }) => <Button variant="outlined" color="primary" onClick={e => copyToClipboard(record[source])}>COPY TO CLIPBOARD</Button>;

export const ProjectList = props => (
    <List {...props} bulkActionButtons={false}>
        <Datagrid>
            <NumberField source="id" label="ProjectID" />
            <TextField source="pilgrimid" label="PilgrimID" />
            <TextField source="title" label="Title" />
            <TextField source="details" label="Details" />
            <NumberField source="cpu" label="CPU Limits (number)" />
            <NumberField source="memory" label="Memory Limits (bytes)" />
            <NumberField source="storage" label="Storage Requests (bytes)" />
            <CopyKubecfgField source="kube_configuration" label="KUBECONFIG" sortable={false} />
            <DeleteButton />
            </Datagrid>
    </List>
);

export const ProjectCreate = props => (
    <Create {...props}>
        <SimpleForm toolbar={<ProjectCreateToolbar />}>
            <TextInput source="title" label="Title" />
            <TextInput source="details" fullWidth='true' label="Details" />
            <NumberInput source="cpu" label="CPU Limits (number)" helperText="Across all pods in a non-terminal state, the sum of CPU limits cannot exceed this value." />
            <NumberInput source="memory" label="Memory Limits (bytes)" helperText="Across all pods in a non-terminal state, the sum of memory limits cannot exceed this value." />
            <NumberInput source="storage" label="Storage Requests (bytes)" helperText="Across all persistent volume claims, the sum of storage requests cannot exceed this value." />
        </SimpleForm>
    </Create>
);

const ProjectCreateToolbar = props => (
    <Toolbar {...props}>
        <SaveButton
            transform={data => ({ ...data })}
        />
    </Toolbar>
);