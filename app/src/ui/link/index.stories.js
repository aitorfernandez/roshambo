import React from 'react'
import StoryRouter from 'storybook-react-router'

import { Link } from '.'

export default {
  component: Link,
  decorators: [StoryRouter()],
  title: 'ui/Link',
}

export const Default = () => (
  <Link to="/">
    {`I'm a Link`}
  </Link>
)
