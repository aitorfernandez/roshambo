import { routes } from './routes'

describe('routes', () => {
  let homePage

  beforeAll(() => {
    ([homePage] = routes)
  })

  describe('homePage', () => {
    it('path', () => {
      expect(homePage.path).toStrictEqual('/')
    })

    it('exact', () => {
      expect(homePage.exact).toBe(true)
    })

    it('component', () => {
      expect(homePage.component.name).toBe('HomePage')
    })
  })
})
