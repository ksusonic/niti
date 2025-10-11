'use client';

import { Section, Cell, List } from '@telegram-apps/telegram-ui';

import { Link } from '@/components/Link/Link';
import { Page } from '@/components/Page';

export default function Home() {


  return (
    <Page back={false}>
      <List>
        <Section
          header="Application Launch Data"
          footer="These pages help developer to learn more about current launch information"
        >
          <Link href="/init-data">
            <Cell subtitle="User data, chat information, technical data">
              Init Data
            </Cell>
          </Link>
          <Link href="/launch-params">
            <Cell subtitle="Platform identifier, Mini Apps version, etc.">
              Launch Parameters
            </Cell>
          </Link>
          <Link href="/theme-params">
            <Cell subtitle="Telegram application palette information">
              Theme Parameters
            </Cell>
          </Link>
        </Section>
      </List>
    </Page>
  );
}
