import { Box, Button, Grid, Group, LoadingOverlay, Modal, PasswordInput, Select, TextInput } from "@mantine/core";
import React, { useEffect, useState } from "react";
import { useForm } from '@mantine/form';
import { usePostUserMutation, useUpdateUserMutation } from "@/redux/services/user";
import { User } from "@/types/users";
import { notifications } from "@mantine/notifications";
import { useCreateUserPermitRoleQuery } from "@/redux/services/role";
import { Role } from "@/types/role";

interface CreateOrUpdateDataInterFace {
  editData: User | null | undefined;
  // eslint-disable-next-line @typescript-eslint/ban-types
  setEditData: Function;
}

interface CreateOrUpdateFormInterFace {
  username: string;
  first_name: string;
  last_name: string;
  email: string;
  phone_number: string;
  role_id: string;
}

const CreateOrUpdateData: React.FC<CreateOrUpdateDataInterFace> = ({ editData, setEditData }) => {


  const [open, setOpen] = useState(false);
  const { data: roles } = useCreateUserPermitRoleQuery();
  console.log({ roles })
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
        <Button onClick={() => { setOpen(true) }} className=" bg-orange-600 hover:bg-orange-700"> Create Company</Button>
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
      form.setValues((prev) => ({ ...prev, ...editData, role_id: String(editData.role_id) }));
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const handleFormSubmit = (values: CreateOrUpdateFormInterFace): void => {

    const customizeValues = { ...values, role_id: Number(values.role_id) }
    if (editData) {
      updateUser({ user_id: editData.id, body: customizeValues })
        .unwrap()
        .then(() => {
          notifications.show({ message: 'Sucessfully Updated' });
          handleModalClose();
        })
        .catch((error: { data: string }) => { notifications.show({ message: error.data, color: 'red' }) })
    }
    else {
      createUser(customizeValues).unwrap()
        .then(() => {
          notifications.show({ message: 'Sucessfully Created' });
          handleModalClose();
        })
        .catch((error: { data: string }) => { notifications.show({ message: error.data, color: 'red' }) })
    }
  }

  return (
    <Box maw={500} mx="auto" pos="relative" px={40}>
      <LoadingOverlay visible={isCreateLoading || isUpdateLoading} overlayBlur={2} />
      <form onSubmit={form.onSubmit((values: CreateOrUpdateFormInterFace): void => handleFormSubmit(values))} >


        <TextInput
          withAsterisk
          label="Name"
          placeholder="your company name..."
          {...form.getInputProps('name')}
        />
        <TextInput
          withAsterisk
          label="Damain"
          placeholder="your domain url..."
          {...form.getInputProps('domain')}
        />

        <Group position="right" mt="md">
          <Button type="submit" variant="filled" className=" bg-orange-600 hover:bg-orange-700" >Submit</Button>
        </Group>
      </form>
    </Box>
  );
}

export default CreateOrUpdateData;