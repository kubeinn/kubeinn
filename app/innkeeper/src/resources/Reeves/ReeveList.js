import * as React from "react";
import {
    List,
    Edit,
    Datagrid,
    TextField,
    EmailField,
    SimpleForm,
    EditButton,
    RadioButtonGroupInput,
} from 'react-admin';

export const ReeveList = props => (
    <List {...props} >
        <Datagrid>
            <TextField source="id" label="ReeveID" />
            <TextField source="villageid" label="VillageID" />
            <TextField source="username" label="Username" />
            <EmailField source="email" label="Email" />
            <TextField source="status" label="Status" />
            <EditButton />
        </Datagrid>
    </List>
);

// export const ReeveCreate = props => (
//     <Create {...props}>
//         <SimpleForm>
//             <TextInput source="villageid" label="VillageID" />
//             <TextInput source="username" label="Username" />
//             <TextInput source="email" label="Email" />
//             <RadioButtonGroupInput source="status" choices={[
//                 { id: 'accepted', name: 'accepted' },
//                 { id: 'pending', name: 'pending' },
//                 { id: 'rejected', name: 'rejected' },
//             ]} />
//         </SimpleForm>
//     </Create>
// );

export const ReeveEdit = props => (
    <Edit {...props}>
        <SimpleForm>
            <RadioButtonGroupInput source="status" choices={[
                { id: 'accepted', name: 'accepted' },
                { id: 'pending', name: 'pending' },
                { id: 'rejected', name: 'rejected' },
            ]} />
        </SimpleForm>
    </Edit>
);