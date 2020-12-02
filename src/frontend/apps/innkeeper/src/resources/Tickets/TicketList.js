import * as React from "react";
import {
    List,
    Edit,
    Create,
    Datagrid,
    TextField,
    EmailField,
    EditButton,
    SimpleForm,
    TextInput,
    RadioButtonGroupInput,
    DeleteButton,
    useNotify,
    useRefresh,
    useRedirect,
} from 'react-admin';

export const TicketList = props => (
    <List {...props} bulkActionButtons={false} >
        <Datagrid>
            <TextField source="id" label="TicketID" />
            <TextField source="pilgrimid" label="PilgrimID" />
            <EmailField source="email" label="Email" />
            <TextField source="topic" label="Topic" />
            <TextField source="details" label="Details" />
            <TextField source="status" label="Status" />
            <EditButton />
            <DeleteButton />
        </Datagrid>
    </List>
);

export const TicketCreate = props => {
    const notify = useNotify();
    const refresh = useRefresh();
    const redirect = useRedirect();
    const onCreateSuccess = ({ data }) => {
        notify("Element created.")
        redirect('/tickets');
        refresh();
    };
    return (
        <Create {...props} onSuccess={onCreateSuccess}>
            <SimpleForm >
                <TextInput source="pilgrimid" label="PilgrimID" />
                <TextInput source="email" label="Email" />
                <TextInput source="topic" label="Topic" />
                <TextInput source="details" label="Details" />
                <RadioButtonGroupInput source="status" label="Status" choices={[
                    { id: 'Open', name: 'Open' },
                    { id: 'Closed', name: 'Closed' },
                ]} />
            </SimpleForm>
        </Create>
    )
};

export const TicketEdit = props => (
    <Edit {...props}>
        <SimpleForm>
            <TextInput source="email" label="Email" />
            <TextInput source="topic" label="Topic" />
            <TextInput source="details" label="Details" />
            <RadioButtonGroupInput source="status" label="Status" choices={[
                { id: 'Open', name: 'Open' },
                { id: 'Closed', name: 'Closed' },
            ]} />
        </SimpleForm>
    </Edit>
);