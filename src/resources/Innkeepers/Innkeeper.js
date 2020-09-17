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
    Filter,
} from 'react-admin';

export const InnkeeperEdit = props => (
    <Edit {...props}>
        <SimpleForm>
            <TextInput disabled source="id" />
            <ReferenceInput source="id" reference="id">
                <SelectInput optionText="name" />
            </ReferenceInput>
            <TextInput source="username" />
            <TextInput source="email" />
        </SimpleForm>
    </Edit>
);

export const InnkeeperList = props => (
    <List {...props} >
        <Datagrid>
            <NumberField source="id" />
            <TextField source="username" />
            <EmailField source="email" />
            <EditButton />
        </Datagrid>
    </List>
);