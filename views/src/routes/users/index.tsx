
import { Table } from "@mantine/core";
import { useLoaderData } from "react-router-dom";

export default function UserIndex() {
  const { users: any } = useLoaderData();

  const columns = [
    {
      title: "Username",
      dataIndex: "username",
      key: "username"
    }
  ]
  return (
    <>
      <Table verticalSpacing="xs" fontSize="xs">
        <thead>
          <tr>
            <td>Username</td>
          </tr>
        </thead>

        <tbody>
          {
            users?.map((user: unknown, index: number) => (<td key={index}>{user.username}</td>))
          }
        </tbody>
      </Table>
      {/* <Table columns={columns} dataSource={users}/> */}
    </>
  )
}


