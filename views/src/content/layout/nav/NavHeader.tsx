import {
  Avatar, Code, Group, Tooltip,
  // rem
} from "@mantine/core"
import { IconLogout } from "@tabler/icons-react"
// import { Logo } from "./Logo"

const NavHeader = () => {
  return (
    <>
      <Group position="apart">
        {/* <Logo clipRule="white" width={rem(120)} /> */}
        <Avatar color="orange" w={100} radius="md"><span className=" text-red-700" >Elit</span>Buzz</Avatar>
        <Code sx={{ fontWeight: 700 }}>v.1</Code>

        <Tooltip label="Log out">
        <Avatar color="orange" translate="yes"  sx={{ transition:"all", ":hover":{cursor:"pointer",scale:1.5}}}>
          <IconLogout />
        </Avatar>
        </Tooltip>
      </Group>
    </>
  )
}

export default NavHeader