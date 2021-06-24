module.exports = {
  collectCoverageFrom: [
    '**/*.{js}',
    '!**/node_modules/**',
  ],
  setupFilesAfterEnv: [
    '<rootDir>/jest/setup.js',
  ],
  testPathIgnorePatterns: [
    '/node_modules/',
  ],
  transformIgnorePatterns: [
    '/node_modules/',
  ],
}
