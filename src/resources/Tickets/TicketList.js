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
    NumberInput,
    SelectInput,
    SimpleForm,
    TextInput,
    RadioButtonGroupInput,
    Filter,
} from 'react-admin';

export const TicketList = props => (
    <List {...props} >
        <Datagrid>
            <TextField source="id" label="TicketID" />
            <TextField source="villageid" label="VillageID" />
            <EmailField source="email" label="Email" />
            <TextField source="topic" label="Topic" />
            <TextField source="details" label="Details" />
            <TextField source="status" label="Status" />
            <EditButton />
        </Datagrid>
    </List>
);

export const TicketCreate = props => (
    <Create {...props}>
        <SimpleForm>
            <TextInput source="villageid" label="VillageID" />
            <TextInput source="email" label="Email" />
            <TextInput source="topic" label="Topic" />
            <TextInput source="details" label="Details" />
            <RadioButtonGroupInput source="status" label="Status" choices={[
                { id: 'Open', name: 'Open' },
                { id: 'Closed', name: 'Closed' },
            ]} />
        </SimpleForm>
    </Create>
);

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