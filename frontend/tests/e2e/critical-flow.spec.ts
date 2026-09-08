import { test, expect } from '@playwright/test';

// Critical E2E flow (theo SKILL.md):
// Register -> Login -> Dashboard -> Course -> Lesson -> Exercise -> Submit -> Result -> XP -> Progress
test('trang chủ hiển thị đúng tiêu đề', async ({ page }) => {
  await page.goto('/');
  await expect(page.getByRole('heading', { name: 'Learning Platform' })).toBeVisible();
});

test.skip('luồng đầy đủ: đăng ký -> học -> nộp bài -> nhận XP', async ({ page }) => {
  // TODO: triển khai khi backend/API sẵn sàng.
});
