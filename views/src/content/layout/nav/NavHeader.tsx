import { Code, Group, rem } from "@mantine/core"
import { Logo } from "./Logo"

const NavHeader = () => {
  return (
    <>
      <Group position="apart">
        <Logo clipRule="white" width={rem(120)} />
        <Code sx={{ fontWeight: 700 }}>v.1</Code>
      </Group>
    </>
  )
}

export default NavHeader