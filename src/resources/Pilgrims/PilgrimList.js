import * as React from "react";
import {
    Show,
    ShowButton,
    SimpleShowLayout,
    RichTextField,
    DateField,
    Pagination,
    useListContext,
    List,
    Edit,
    Create,
    Datagrid,
    ReferenceField,
    TextField,
    EmailField,
    NumberField,
    EditButton,
    ReferenceInput,
    SelectInput,
    SimpleForm,
    TextInput,
    PasswordInput,
    Filter,
} from 'react-admin';

export const PilgrimList = props => (
    <List {...props} >
        <Datagrid>
            <NumberField source="id" />
            <TextField source="username" />
            <EmailField source="email" />
        </Datagrid>
    </List>
);

export const PilgrimCreate = props => (
    <Create {...props}>
        <SimpleForm>
            <TextInput source="username" />
            <TextInput source="email" />
            <PasswordInput label="Password" source="passwd" helperText="Default password: pilgrim" initiallyVisible />
        </SimpleForm>
    </Create>
);