import React from 'react'
import { ThemeProvider } from 'styled-components'

import { AppGlobalStyle } from '../src/AppGlobalStyle'
import { theme } from '../src/theme'

export const parameters = {
  actions: { argTypesRegex: "^on[A-Z].*" },
}

export const decorators = [
  (Story) => (
    <ThemeProvider theme={theme}>
      <AppGlobalStyle />
      <Story />
    </ThemeProvider>
  )
]
