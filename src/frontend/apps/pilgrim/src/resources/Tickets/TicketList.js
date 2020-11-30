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
    useNotify,
    useRefresh,
    useRedirect,
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
        <Create {...props}>
            <SimpleForm onSuccess={onCreateSuccess}>
                <TextInput source="email" label="Email" />
                <TextInput source="topic" label="Topic" />
                <TextInput source="details" label="Details" />
            </SimpleForm>
        </Create>
    )
};