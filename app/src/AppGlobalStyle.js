import { createGlobalStyle } from 'styled-components'

export const AppGlobalStyle = createGlobalStyle`
  *,
  *::before,
  *::after {
    box-sizing: border-box;
  }

  html {
    -ms-text-size-adjust: 100%;
    -webkit-text-size-adjust: 100%;
    font-family: sans-serif;
  }

  html, body {
    height: 100%;
  }

  body {
    -moz-font-feature-settings: 'liga' on;
    -moz-osx-font-smoothing: grayscale;
    -webkit-font-smoothing: antialiased;
    line-height: 1.5;
    margin: 0;
    text-rendering: optimizeLegibility;
    word-wrap: break-word;
  }
`
