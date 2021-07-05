const borderRadius = [
  '0',
  '2px',
  '4px',
  '6px',
  '12px',
  '24px',
]

const color = {
  gray: [
    '#fafbfc',
    '#f6f8fa',
    '#e1e4e8',
    '#d1d5da',
    '#959da5',
    '#6a737d',
    '#586069',
    '#444d56',
    '#2f363d',
    '#24292e',
  ],
  red: [
    '#ffeef0',
    '#ffdce0',
    '#fdaeb7',
    '#f97583',
    '#ea4a5a',
    '#d73a49',
    '#cb2431',
    '#b31d28',
    '#9e1c23',
    '#86181d',
  ],
  green: [
    '#f0fff4',
    '#dcffe4',
    '#bef5cb',
    '#85e89d',
    '#34d058',
    '#28a745',
    '#22863a',
    '#176f2c',
    '#165c26',
    '#144620',
  ],
  yellow: [
    '#fffdef',
    '#fffbdd',
    '#fff5b1',
    '#ffea7f',
    '#ffdf5d',
    '#ffd33d',
    '#f9c513',
    '#dbab09',
    '#b08800',
    '#735c0f',
  ],
  black: [
    'rgba(27, 31, 35, 0.15)',
    'rgba(27, 31, 35, 0.3)',
    'rgba(27, 31, 35, 0.5)',
    'rgba(27, 31, 35, 0.7)',
    'rgba(27, 31, 35, 0.85)',
    'rgba(27, 31, 35, 1)',
  ],
  white: [
    'rgba(255, 255, 255, 0.15)',
    'rgba(255, 255, 255, 0.3)',
    'rgba(255, 255, 255, 0.5)',
    'rgba(255, 255, 255, 0.7)',
    'rgba(255, 255, 255, 0.85)',
    'rgba(255, 255, 255, 1)',
  ],
}

function fontStack(fonts) {
  return fonts.map((f) => (f.includes(' ') ? `"${f}"` : f)).join(', ')
}

const fontFamily = {
  primary: fontStack([
    'Open Sans',
    'sans-serif',
  ]),
  secondary: fontStack([
    'Varela Round',
    'sans-serif',
  ]),
}

const fontBaseSize = 16

function fontConverter(sizes) {
  const parse = (v) => parseInt(v, 10)
  return sizes.map((s) => `${(1 / fontBaseSize) * parse(s)}rem`)
}

const fontSize = fontConverter([
  '12px',
  '14px',
  '16px',
  '20px',
  '24px',
  '32px',
  '44px',
  '56px',
  '96px',
  '120px',
])

const fontWeight = {
  light: 300,
  normal: 400,
  bold: 600,
}

const opacity = [
  0, 0.15, 0.3, 0.5, 0.7, 0.85, 1,
]

const size = {
  sm: '544px',
  md: '768px',
  lg: '1012px',
  xl: '1280px',
}

const space = [
  '0',
  '2px',
  '4px',
  '8px',
  '12px',
  '16px', // 5
  '20px',
  '24px',
  '32px',
  '40px',
  '48px', // 10
  '64px',
  '80px',
  '96px',
]

const transition = {
  ease: '.2s ease',
}

export const theme = {
  borderRadius,
  color,
  fontBaseSize,
  fontFamily,
  fontSize,
  fontWeight,
  opacity,
  size,
  space,
  transition,
}
