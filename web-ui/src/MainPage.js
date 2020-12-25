import React from "react";
import {
  PageWithHeader,
  TopNav,
  Box,
  Button,
  Image,
  PageContent,
} from "bumbag";


export default function MainPage() {
  return (
    <PageWithHeader
      sticky
      header={
        <TopNav>
          <TopNav.Section>
            <TopNav.Item href="https://bumbag.style" fontWeight="semibold">
              <Image src="/logo.png" height="44px" />
            </TopNav.Item>
            <TopNav.Item href="#">Get started</TopNav.Item>
            <TopNav.Item href="#">Components</TopNav.Item>
          </TopNav.Section>
          <TopNav.Section marginRight="major-2">
            <TopNav.Item>
              <Button variant="ghost" palette="primary">
                Sign up
              </Button>
            </TopNav.Item>
            <TopNav.Item>
              <Button palette="primary">Login</Button>
            </TopNav.Item>
          </TopNav.Section>
        </TopNav>
      }
      border="default"
      overrides={{
        PageWithHeader: { styles: { base: { minHeight: "unset" } } },
      }}
    >
      <Box>
        <PageContent isFluid wrapperProps={{ backgroundColor: "white900" }}>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
          eiusmod tempor incididunt ut labore et dolore magna aliqua. Arcu
          bibendum at varius vel. Volutpat sed cras ornare arcu dui. Faucibus
          scelerisque eleifend donec pretium vulputate sapien nec. Faucibus
          pulvinar elementum integer enim neque volutpat. Cum sociis natoque
          penatibus et magnis dis parturient montes. Maecenas accumsan lacus vel
          facilisis. Mauris pellentesque pulvinar pellentesque habitant morbi
          tristique senectus. Gravida cum sociis natoque penatibus et magnis dis
          parturient montes. Massa sapien faucibus et molestie ac feugiat sed
          lectus vestibulum. Nisi est sit amet facilisis magna etiam tempor.
          Eget nulla facilisi etiam dignissim diam quis enim lobortis
          scelerisque.
        </PageContent>
      </Box>
    </PageWithHeader>
  );
}
