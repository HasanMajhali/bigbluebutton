const { test } = require('../fixtures');
const { SharedNotes } = require('./sharednotes');
const { initializePages } = require('../core/helpers');
const { fullyParallel } = require('../playwright.config');

test.describe('Shared Notes', { tag: ['@ci', '@flaky-3.1'] }, () => {
  const sharedNotes = new SharedNotes();

  test.describe.configure({ mode: fullyParallel ? 'parallel' : 'serial' });
  test[fullyParallel ? 'beforeEach' : 'beforeAll'](async ({ browser }) => {
    await initializePages(sharedNotes, browser, { isMultiUser: true });
  });

  test('Open shared notes', async () => {
    await sharedNotes.openSharedNotes();
  });

  test('Type in shared notes', async () => {
    await sharedNotes.typeInSharedNotes();
  });

  test('Formate text in shared notes', async () => {
    await sharedNotes.formatTextInSharedNotes();
  });

  test('Export shared notes', async ({}, testInfo) => {
    await sharedNotes.exportSharedNotes(testInfo);
  });

  test('Convert notes to presentation', async () => {
    await sharedNotes.convertNotesToWhiteboard();
  });

  test('Multiusers edit', async () => {
    await sharedNotes.editSharedNotesWithMoreThanOneUSer();
  });

  test('See notes without edit permission', async () => {
    await sharedNotes.seeNotesWithoutEditPermission();
  });

  // different failures in CI and local
  // local: not able to click on "unpin" button
  // CI: not restoring presentation for viewer after unpinning notes
  test('Pin and unpin notes onto whiteboard', { tag: '@flaky' }, async () => {
    await sharedNotes.pinAndUnpinNotesOntoWhiteboard();
  });
});
