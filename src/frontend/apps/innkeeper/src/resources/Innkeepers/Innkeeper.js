import * as React from "react";
import {
    List,
    Datagrid,
    TextField,
    EmailField,
    DeleteButton,
    EditButton,
    TextInput,
    Create,
    SimpleForm,
    PasswordInput,
    Edit,
    useNotify,
    useRefresh,
    useRedirect,
} from 'react-admin';

export const InnkeeperList = props => (
    <List {...props} bulkActionButtons={false} >
        <Datagrid>
            <TextField source="id" label="InnkeeperID" />
            <TextField source="username" label="Username" />
            <EmailField source="email" label="Email" />
            <TextField source="passwd" label="Password" />
            <EditButton />
            <DeleteButton />
        </Datagrid>
    </List>
);


export const InnkeeperCreate = props => {
    const notify = useNotify();
    const refresh = useRefresh();
    const redirect = useRedirect();
    const onCreateSuccess = ({ data }) => {
        notify("Element created.")
        redirect('/innkeepers');
        refresh();
    };
    return (
        <Create {...props} onSuccess={onCreateSuccess}>
            <SimpleForm>
                <TextInput source="username" label="Username" />
                <TextInput source="email" label="Email" />
                <PasswordInput label="Password" source="passwd" />
            </SimpleForm>
        </Create>
    )
};

export const InnkeeperEdit = (props) => {
    return (
        <Edit {...props} >
            <SimpleForm>
                <TextInput source="username" label="Username" />
                <TextInput source="email" label="Email" />
                <PasswordInput label="Password" source="passwd" />
            </SimpleForm>
        </Edit>
    );
}