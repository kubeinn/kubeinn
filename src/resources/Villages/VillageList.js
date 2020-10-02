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

export const VillageList = props => (
    <List {...props} >
        <Datagrid>
            <NumberField source="id" />
            <TextField source="title" />
            <TextField source="details" />
        </Datagrid>
    </List>
);

export const VillageCreate = props => (
    <Create {...props}>
        <SimpleForm>
            <TextInput source="title" />
            <TextInput source="details" fullWidth='true'/>
        </SimpleForm>
    </Create>
);