import * as React from "react";
import {
    List,
    Datagrid,
    TextField,
    EmailField,
    Create,
    SimpleForm,
    TextInput,
    Edit,
    RadioButtonGroupInput,
    useNotify,
    useRefresh,
    useRedirect,
    EditButton,
    DeleteButton,
} from 'react-admin';

export const PilgrimList = props => (
    <List {...props} bulkActionButtons={false} >
        <Datagrid>
            <TextField source="id" label="PilgrimID" />
            <TextField source="organization" label="Organization" />
            <TextField source="description" label="Description" fullWidth />
            <TextField source="username" label="Username" />
            <TextField source="email" label="Email" />
            <TextField source="passwd" label="Password" />
            <TextField source="status" label="Status" />
            <EditButton />
            <DeleteButton />
        </Datagrid>
    </List>
);

export const PilgrimCreate = props => {
    const notify = useNotify();
    const refresh = useRefresh();
    const redirect = useRedirect();
    const onCreateSuccess = ({ data }) => {
        notify("Element created.")
        redirect('/pilgrims');
        refresh();
    };
    return (
        <Create {...props} onSuccess={onCreateSuccess}>
            <SimpleForm>
                <TextInput source="organization" label="Organization" />
                <TextInput source="description" label="Description" fullWidth />
                <TextInput source="username" label="Username" />
                <TextInput source="email" label="Email" />
                <TextInput source="passwd" label="Password" />
                <RadioButtonGroupInput source="status" choices={[
                    { id: 'accepted', name: 'accepted' },
                    { id: 'pending', name: 'pending' },
                    { id: 'rejected', name: 'rejected' },
                ]} />
            </SimpleForm>
        </Create>
    )
};

export const PilgrimEdit = (props) => {
    return (
        <Edit {...props} >
            <SimpleForm>
                <TextInput source="organization" label="Organization" />
                <TextInput source="description" label="Description" fullWidth />
                <TextInput source="username" label="Username" />
                <TextInput source="email" label="Email" />
                <TextInput source="passwd" label="Password" />
                <RadioButtonGroupInput source="status" choices={[
                    { id: 'accepted', name: 'accepted' },
                    { id: 'pending', name: 'pending' },
                    { id: 'rejected', name: 'rejected' },
                ]} />
            </SimpleForm>
        </Edit>
    );
}