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
    RadioButtonGroupInput,
    Filter,
} from 'react-admin';

export const TicketList = props => (
    <List {...props} >
        <Datagrid>
            <NumberField source="id" />
            <TextField source="pilgrimID" />
            <EmailField source="email" />
            <TextField source="topic" />
            <TextField source="details" />
            <TextField source="status" />
        </Datagrid>
    </List>
);

export const TicketCreate = props => (
    <Create {...props}>
        <SimpleForm>
            <TextInput source="email" />
            <TextInput source="topic" />
            <TextInput source="details" fullWidth='true' />
            <RadioButtonGroupInput source="status" choices={[
                { id: 'Open', name: 'Open' },
                { id: 'Closed', name: 'Closed' },
            ]} />
        </SimpleForm>
    </Create>
);