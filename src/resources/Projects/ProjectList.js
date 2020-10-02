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
    NumberInput,
    Filter,
} from 'react-admin';

export const ProjectList = props => (
    <List {...props} >
        <Datagrid>
            <NumberField source="id" />
            <TextField source="title" />
            <TextField source="details" />
            <NumberField source="cpu" />
            <NumberField source="memory" />
            <NumberField source="storage" />
        </Datagrid>
    </List>
);

export const ProjectCreate = props => (
    <Create {...props}>
        <SimpleForm>
            <TextInput source="title" />
            <TextInput source="details" fullWidth='true' />
            <NumberInput source="cpu" />
            <NumberInput source="memory" />
            <NumberInput source="storage" />
        </SimpleForm>
    </Create>
);