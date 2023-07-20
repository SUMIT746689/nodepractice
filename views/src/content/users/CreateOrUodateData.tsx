import { Box, Button, Grid, Group, Modal, Select, TextInput } from "@mantine/core";
import React, { useEffect, useState } from "react";
import { useForm } from '@mantine/form';
import { usePostUserMutation } from "@/redux/services/user";
import { User } from "@/types/users";

interface CreateOrUodateDataInterFace {
  editData: User | null | undefined;
  // eslint-disable-next-line @typescript-eslint/ban-types
  setEditData: Function;
}

const CreateOrUpdateData: React.FC<CreateOrUodateDataInterFace> = ({ editData,setEditData }) => {
  const [open, setOpen] = useState(false);
  useEffect(() => {
    if (editData) setOpen(true);
  }, [editData]);

  const handleModalClose = ()=>{
    setOpen(false);
    setEditData(undefined)
  }
  return (
    <>
      <Modal opened={open} onClose={handleModalClose} title="Authentication" centered>
        {/* Modal content */}
        <Form editData={editData} />
      </Modal>

      <Group position="center" pr={20} >
        <Button onClick={() => { setOpen(true) }} className=" bg-orange-600 hover:bg-orange-700"> Create User</Button>
      </Group>
    </>
  )
}

interface FormInterface {
  editData: User | undefined | null;
}

const Form: React.FC<FormInterface> = ({ editData }) => {
  const [updateUser] = usePostUserMutation()

  const form = useForm({
    initialValues: {
      first_name: '',
      last_name: '',
      phone_number: '',
      email: '',
      username: '',
      password: '',
      confirm_password: '',
      role: '',
      // termsOfService: false,
    },
    validate: {
      email: (value) => (/^\S+@\S+$/.test(value) ? null : 'Invalid email'),
    },
  });
  
  let role = "ADMIN"
  const super_admin_roles = ['ADMIN']
  const admin_roles = ['CASHIER', 'CUSTOMER', 'SUPPLIER', 'VENDOR']
  const cashier_roles = ['CUSTOMER', 'SUPPLIER', 'VENDOR']


  useEffect(() => {
    if (editData) {
      form.setValues((prev) => ({ ...prev, ...editData }));
      // return form.reset();
    }
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  console.log({ form: form.values })
  const handleFormSubmit = async (values:User) => {
    console.log({ values })
    const { data,error } = await updateUser(values);
    console.log({ data,error })
  }

  return (
    <Box maw={500} mx="auto" px={40}>
      <form onSubmit={form.onSubmit(handleFormSubmit)}>

        <Grid grow gutter="xs">
          <Grid.Col span={4}>
            <TextInput
              withAsterisk
              label="First name"
              placeholder="your first name..."
              {...form.getInputProps('first_name')}
            />
          </Grid.Col>
          <Grid.Col span={4}>
            <TextInput
              withAsterisk
              label="Last name"
              placeholder="your last name..."
              {...form.getInputProps('last_name')}
            />
          </Grid.Col>
        </Grid>

        <TextInput
          // withAsterisk
          label="Phone Number"
          placeholder="your phone number..."
          {...form.getInputProps('phone_number')}
        />

        <TextInput
          // withAsterisk
          label="Email"
          placeholder="your@email.com"
          {...form.getInputProps('email')}
        />
        <TextInput
          withAsterisk
          label="Username"
          placeholder="your username.."
          {...form.getInputProps('username')}
        />
        <TextInput
          withAsterisk
          label="Password"
          placeholder="your password..."
          {...form.getInputProps('password')}
        />
        <TextInput
          withAsterisk
          label="Confirm Password"
          placeholder="your confirm password..."
          {...form.getInputProps('confirm_password')}
          />

        <Select
          label="Select Role"
          placeholder="Select Role"
          {...form.getInputProps('role')}
          // sx={{"::selection":{backgroundColor:'orange'}}}
          data={
            role === "SUPERADMIN" && super_admin_roles.map(value =>({value,label:value})) || 
            role === "ADMIN" && admin_roles.map(value =>({value,label:value})) || 
            role === "TEACHER" && cashier_roles.map(value =>({value,label:value})) || 
            []
          }
        />


        {/* <Checkbox
          mt="md"
          label="I agree to sell my privacy"
          {...form.getInputProps('termsOfService', { type: 'checkbox' })}
        /> */}

        <Group position="right" mt="md">
          <Button type="submit" variant="filled" className=" bg-orange-600 hover:bg-orange-700" >Submit</Button>
        </Group>
      </form>
    </Box>
  );
}

export default CreateOrUpdateData;