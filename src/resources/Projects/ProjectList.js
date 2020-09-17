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