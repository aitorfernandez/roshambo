export function readQuery(cache, query, vars) {
  return cache.readQuery({
    query,
    variables: {
      ...vars,
    },
  })
}

export function writeQuery(cache, query, data, vars) {
  cache.writeQuery({
    query,
    data,
    variables: {
      ...vars,
    },
  })
}
