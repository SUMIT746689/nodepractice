import { Box, Button, Grid, Group, LoadingOverlay, Modal, Select, TextInput } from "@mantine/core";
import React, { useEffect, useState } from "react";
import { useForm } from '@mantine/form';
import { usePostUserMutation, useUpdateUserMutation } from "@/redux/services/user";
import { UpdateUserBody, User } from "@/types/users";
import { AuthUser } from "@/types/auth";
import { notifications } from "@mantine/notifications";

interface CreateOrUodateDataInterFace {
  editData: User | null | undefined;
  // eslint-disable-next-line @typescript-eslint/ban-types
  setEditData: Function;
  // eslint-disable-next-line @typescript-eslint/ban-types
  // refetch: ()=>void;
  authUser: AuthUser | undefined;
}

const CreateOrUpdateData: React.FC<CreateOrUodateDataInterFace> = ({ editData, setEditData, authUser }) => {
  const [open, setOpen] = useState(false);

  useEffect(() => {
    if (editData) setOpen(true);
  }, [editData]);

  const handleModalClose = () => {
    setOpen(false);
    setEditData(undefined)
  }
  return (
    <>
      <Modal opened={open} onClose={handleModalClose} title="User" centered>
        {/* Modal content */}
        <Form editData={editData} handleModalClose={handleModalClose} role={authUser?.role} />
      </Modal>

      <Group position="center" pr={20} >
        <Button onClick={() => { setOpen(true) }} className=" bg-orange-600 hover:bg-orange-700"> Create User</Button>
      </Group>
    </>
  )
}

interface FormInterface {
  editData: User | undefined | null;
  // refetch: () => void;
  handleModalClose: () => void;
  role: string | undefined;
}

const Form: React.FC<FormInterface> = ({ editData, handleModalClose, role }) => {
  const [createUser, { isLoading: isCreateLoading }] = usePostUserMutation()
  const [updateUser, { isLoading: isUpdateLoading }] = useUpdateUserMutation()


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

  const handleFormSubmit = async (values: UpdateUserBody): Promise<void> => {
    console.log({ values })
    try {
      if (editData) {
        const { data: ud, error } = await updateUser({ user_id: editData.id, body: values });
        console.log({ ud, error });
        // if (!error?.data) throw new Error(error.message);
        if (error) throw new Error(error.data);
        notifications.show({ message: 'Sucessfully Updated' });

      }
      else {
        const { data: cd, error } = await createUser(values);
        console.log({ data: cd, error })
        // if (!error?.data) throw new Error(error.message);
        if (error) throw new Error(error.data);
        notifications.show({ message: 'Sucessfully Created' });
      }

      // refetch();
      handleModalClose();
    } catch (err) {
      console.log({ err });
      notifications.show({ message: err.message, color: 'red' });
    }

  }

  return (
    <Box maw={500} mx="auto" pos="relative" px={40}>
      <LoadingOverlay visible={isCreateLoading || isUpdateLoading} overlayBlur={2} />
      <form onSubmit={form.onSubmit((values: UpdateUserBody): object => handleFormSubmit(values))} >

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
            role === "SUPERADMIN" && super_admin_roles.map(value => ({ value, label: value })) ||
            role === "ADMIN" && admin_roles.map(value => ({ value, label: value })) ||
            role === "TEACHER" && cashier_roles.map(value => ({ value, label: value })) ||
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