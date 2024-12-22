function getValue(obj, key) {
  try {
    key = key && key.split ? key.split('.') : [key]
    for (let i = 0; i < key.length; i += 1) {
      obj = obj[key[i]]
    }
    return obj
  } catch (e) {
    console.error(e, key)
  }
}

export const get = (path) => (props) => {
  return getValue(props.theme, path)
}

export const createMediaQuery = (getFunc) => (props) => {
  const v = getFunc(props)
  return `@media screen and (min-width: ${v})`
}
