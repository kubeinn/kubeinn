import * as React from "react";
import {
    List,
    Create,
    Datagrid,
    TextField,
    SimpleForm,
    TextInput,
} from 'react-admin';

export const VillageList = props => (
    <List {...props} >
        <Datagrid>
            <TextField source="id" label="VillageID" />
            <TextField source="organization" label="Organization"  />
            <TextField source="description" label="Description" />
        </Datagrid>
    </List>
);

export const VillageCreate = props => (
    <Create {...props}>
        <SimpleForm>
            <TextInput source="organization" label="Organization" />
            <TextInput source="description" multiline fullWidth='true' label="Description"/>
        </SimpleForm>
    </Create>
);

// export const VillageEdit = props => (
//     <Edit {...props}>
//         <SimpleForm>
//             <TextInput source="description" multiline fullWidth='true' label="Description" />
//         </SimpleForm>
//     </Edit>
// );