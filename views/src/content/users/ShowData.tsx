export default function ShowData() {
  return (
    <>
      <div className="px-6 py-3 w-full min-h-fit">
        <PaginationExampleWithPageSizeSelector
          headColumns={[
            { accessor: 'id', width: 100 },
            { accessor: 'first_name' },
            { accessor: 'last_name' },
            { accessor: 'username' },
            { accessor: 'email' },
            { accessor: 'phone_number' },
            {
              accessor: 'actions',
              title: <Text mr="xs">Row actions</Text>,
              textAlignment: 'right',
              render: (data) => (
                <Group spacing={4} position="right" noWrap>
                  {/* <ActionIcon color="green" onClick={() => showInfo(data)}>
                    <IconEye size={16} />
                  </ActionIcon> */}
                  <ActionIcon color="blue" onClick={() => setEditUser(data)}>
                    <IconEdit size={16} />
                  </ActionIcon>
                  {/* <ActionIcon color="red" onClick={() => deleteCompany(data)}>
                    <IconTrash size={16} />
                  </ActionIcon> */}
                </Group>
              ),
            },
          ]}
          datas={users?.users || []}
        />
      </div>
    </>
  )
}