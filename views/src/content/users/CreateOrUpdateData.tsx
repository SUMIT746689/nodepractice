import { Box, Button, Grid, Group, LoadingOverlay, Modal, PasswordInput, Select, TextInput } from "@mantine/core";
import React, { useEffect, useState } from "react";
import { useForm } from '@mantine/form';
import { usePostUserMutation, useUpdateUserMutation } from "@/redux/services/user";
import { CreateUser, UpdateUserBody, User } from "@/types/users";
import { AuthUser } from "@/types/auth";
import { notifications } from "@mantine/notifications";
import { useCreateUserPermitRoleQuery } from "@/redux/services/role";
import { Role } from "@/types/role";

interface CreateOrUodateDataInterFace {
  editData: User | null | undefined;
  // eslint-disable-next-line @typescript-eslint/ban-types
  setEditData: Function;
  authUser: AuthUser | undefined;
}

const CreateOrUpdateData: React.FC<CreateOrUodateDataInterFace> = ({ editData, setEditData, authUser }) => {

  console.log({ authUser })
  const [open, setOpen] = useState(false);
  const { data: roles } = useCreateUserPermitRoleQuery();
  // if (authUser) {
  //   const { edges: { role: { edges: {permissions = [] } = undefined } = undefined } = undefined } = authUser;
  //   console.log({ permissions })
  // }

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
        <Form editData={editData} handleModalClose={handleModalClose} roles={roles || []} />
      </Modal>

      <Group position="center" pr={20} >
        <Button onClick={() => { setOpen(true) }} className=" bg-orange-600 hover:bg-orange-700"> Create User</Button>
      </Group>
    </>
  )
}

interface FormInterface {
  editData: User | undefined | null;
  handleModalClose: () => void;
  roles: Role[] | [];
}

const Form: React.FC<FormInterface> = ({ editData, handleModalClose, roles }) => {
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
      role_id: '',
      // termsOfService: false,
    },
    validate: {
      first_name: (value) => (value ? null : 'Type your first name'),
      last_name: (value) => (value ? null : 'Type your last name'),
      phone_number: (value) => (!value || /^\d+$/.test(value) ? null : 'Invalid phone number'),
      email: (value) => (!value || /^\S+@\S+$/.test(value) ? null : 'Invalid email'),
      username: (value) => (value ? null : 'Provide username'),
      password: (value) => (editData || /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$/.test(value) ? null : 'Minimum eight characters, at least one letter and one number'),
      confirm_password: (value, values) => ((value === values.password) ? null : 'Password did not match'),
      role_id: (value) => (/^\d+$/.test(value) ? null : 'Select a role'),
    },
  });

  useEffect(() => {
    if (editData) {
      form.setValues((prev) => ({ ...prev, ...editData,role_id: String(editData.role_id)}));
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const handleFormSubmit = async (values: CreateUser): Promise<void> => {
    console.log({ values })
    const customizeValues = {...values,role_id:Number(values.role_id)}
    try {
      if (editData) {
        const payload = await updateUser({ user_id: editData.id, body: customizeValues });
        console.log({payload});
        // if (!error?.data) throw new Error(error.message);
        // if (error) throw new Error(error.data);
        notifications.show({ message: 'Sucessfully Updated' });

      }
      else {
        const  payload = await createUser(customizeValues);
        console.log({"cr payload": payload})
        // if (!error?.data) throw new Error(error.message);
        // if (error) throw new Error(error.data);
        notifications.show({ message: 'Sucessfully Created' });
      }

      // refetch();
      handleModalClose();
    } catch (err) {
      console.log({ err });
      notifications.show({ message: err.message, color: 'red' });
    }

  }
  console.log({ roles })
  console.log({ editData })

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
          type="number"
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
          placeholder="your username..."
          autoComplete="new-username"
          {...form.getInputProps('username')}
        />
        <PasswordInput
          withAsterisk
          label="Password"
          placeholder="your password..."
          autoComplete="new-password"
          {...form.getInputProps('password')}
        />
        <PasswordInput
          withAsterisk
          label="Confirm Password"
          placeholder="your confirm password..."
          {...form.getInputProps('confirm_password')}
        />

        <Select
          withAsterisk
          label="Select Role"
          placeholder="Select Role"
          {...form.getInputProps('role_id')}
          // sx={{"::selection":{backgroundColor:'orange'}}}
          data={
            roles ? roles?.map(({ id, title }: { id: number, title: string }) => ({ value: String(id), label: title })) : []

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