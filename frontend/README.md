# Learning Platform — Frontend

Frontend cho nền tảng học lập trình tương tác: khoá học, bài học (MDX), bài
tập chấm code, quiz, theo dõi tiến độ/XP/thành tích và khu vực quản trị.

Project được khởi tạo theo skill `frontend-learning-platform` (kiến trúc,
routes, quy tắc coding) kết hợp chuẩn Next.js App Router hiện đại.

## 1. Tech stack

| Nhóm | Công nghệ |
|---|---|
| Framework | Next.js (App Router) + TypeScript + React |
| UI | Tailwind CSS, shadcn/ui, Radix UI, Lucide React |
| Nội dung bài học | MDX (`@next/mdx`) |
| Form & validate | React Hook Form + Zod |
| Code editor | Monaco Editor (lazy-loaded) |
| Chạy code phía client | Pyodide (Python trong browser cho nút "Run") |
| Testing | Vitest + React Testing Library (unit), Playwright (E2E) |
| Theme | next-themes (light/dark) |
| Font | Be Vietnam Pro (`next/font/google`, biến `--font-be-vietnam-pro`) |

Bảng màu và bố cục dashboard (navbar full-width + sidebar 2 cấp thu gọn +
banner/streak/stat card) được thiết kế theo phong cách tham khảo dạng
nền tảng học trực tuyến (tông đỏ chủ đạo, nền kem ấm, card bo tròn lớn) —
xem `app/globals.css` (CSS variables) và `components/layout/sidebar-nav.tsx`
(menu collapsible dùng Radix Collapsible).

## 2. Cấu trúc thư mục

```
learning-platform/
├── app/
│   ├── (public)/            # Trang công khai: /, /courses, /login, /register
│   ├── (student)/           # Layout riêng có sidebar: dashboard, progress, exercises...
│   ├── admin/                # Khu vực quản trị: courses, lessons, submissions, analytics...
│   ├── api/                  # Route handlers (proxy/BFF cho backend nếu cần)
│   ├── layout.tsx            # Root layout + ThemeProvider
│   └── globals.css           # Design tokens (CSS variables) cho Tailwind/shadcn
├── components/
│   ├── ui/                   # Primitives kiểu shadcn (button, card, progress...)
│   ├── layout/                # Sidebar, Navbar, MobileNav
│   ├── mdx/                   # Callout, Tip, Warning, Note, CodeBlock, ExerciseEmbed,
│   │                          # QuizEmbed, Steps, Definition — dùng trong nội dung MDX
│   ├── dashboard/              # StatCard, CourseProgressCard
│   ├── course/                  # CourseCard
│   └── exercise/                 # CodeEditor (Monaco), ExercisePanel (Run/Submit)
├── content/                        # Nội dung MDX theo khoá học
│   ├── python/{basic,advanced}/
│   └── web/{basic,advanced}/
├── lib/
│   ├── api/                         # apiClient + hàm gọi API theo domain (courses, exercises)
│   ├── validations/                  # Zod schema (auth, ...)
│   ├── constants.ts                   # ROUTES tập trung
│   └── utils.ts                        # cn(), formatDate()
├── types/                                # Domain types: course, lesson, exercise, quiz, user, progress
├── hooks/                                 # Custom hooks (trống, bổ sung dần)
├── tests/
│   ├── unit/                              # Vitest + Testing Library
│   └── e2e/                                # Playwright — luồng E2E quan trọng
└── public/
```

### Quy ước route (theo SKILL.md)

- **Public**: `/`, `/courses`, `/courses/[slug]`, `/login`, `/register`
- **Student** (trong nhóm `(student)`, có sidebar): `/dashboard`, `/courses`,
  `/courses/[courseSlug]/[lessonSlug]`, `/exercises/[exerciseId]`,
  `/quizzes/[quizId]`, `/progress`, `/achievements`, `/leaderboard`,
  `/profile`, `/settings`
- **Admin**: `/admin`, `/admin/courses`, `/admin/lessons`,
  `/admin/exercises`, `/admin/test-cases`, `/admin/quizzes`,
  `/admin/submissions`, `/admin/users`, `/admin/analytics`

## 3. Nguyên tắc kiến trúc đã áp dụng

- **Server Components mặc định**: các trang course/lesson/dashboard fetch dữ
  liệu trên server, không có `"use client"` trừ khi thật sự cần tương tác
  (CodeEditor, ExercisePanel, LoginForm, MobileNav, ThemeProvider).
- **Run vs Submit tách biệt**: `ExercisePanel` gọi `exercisesApi.run()` (chạy
  nhanh, phản hồi tức thời) và `exercisesApi.submit()` (gửi backend → queue →
  code runner → kết quả chính thức). Không lộ hidden test case ở frontend.
- **MDX tách khỏi dữ liệu nghiệp vụ**: nội dung bài học nằm ở `content/`,
  còn exercise/quiz/submission/progress/XP là dữ liệu từ backend, nhúng vào
  MDX qua `<ExerciseEmbed>` / `<QuizEmbed>`.
- **API layer tách khỏi UI**: mọi lời gọi API đi qua `lib/api/*`, component
  không tự `fetch()` trực tiếp.
- **Type an toàn**: `types/*.ts` định nghĩa hợp đồng dữ liệu dùng chung giữa
  API layer và component.
- **Lazy-load phần nặng**: Monaco Editor được `dynamic(..., { ssr: false })`.

## 4. Bắt đầu

```bash
npm install
cp .env.example .env.local   # điền NEXT_PUBLIC_API_URL trỏ tới backend
npm run dev                  # http://localhost:3000
```

Các lệnh khác:

```bash
npm run typecheck   # kiểm tra type
npm run lint         # ESLint
npm run test          # Vitest (unit)
npm run test:e2e       # Playwright (cần `npm run dev` hoặc để Playwright tự chạy)
npm run format           # Prettier + sắp xếp class Tailwind
```

> **Lưu ý**: repo này là *scaffold* — `package.json` liệt kê đúng dependency
> cần thiết nhưng **chưa chạy `npm install`** trong quá trình tạo (không có
> mạng/registry npm khi build). Sau khi giải nén, chạy `npm install` trước
> tiên.

## 4.1. Nếu Tailwind không render style (trang hiển thị không có CSS)

Kiểm tra theo thứ tự sau — nguyên nhân phổ biến nhất được liệt kê trước:

1. **Đã chạy `npx shadcn@latest init`?** — bản shadcn CLI mới nhất mặc định
   cài **Tailwind v4** (cú pháp CSS hoàn toàn khác: `@import "tailwindcss";`
   thay vì 3 dòng `@tailwind ...`, không dùng `tailwind.config.ts` theo kiểu
   cũ, đổi plugin PostCSS thành `@tailwindcss/postcss`). Project này được
   viết cho **Tailwind v3**. Nếu đã lỡ chạy init và nó ghi đè
   `postcss.config.mjs` / `globals.css` / cài `tailwindcss@4`, hãy:
   ```bash
   pnpm remove tailwindcss @tailwindcss/postcss
   pnpm add -D tailwindcss@^3.4.0 postcss@^8.4.0 autoprefixer@^10.4.0
   ```
   rồi kiểm tra lại `postcss.config.mjs` đúng như trong repo gốc (plugin
   `tailwindcss` + `autoprefixer`, không phải `@tailwindcss/postcss`) và
   `app/globals.css` bắt đầu bằng 3 dòng `@tailwind base/components/utilities`.
2. **Xoá cache và chạy lại**:
   ```bash
   rm -rf .next
   pnpm dev
   ```
3. **Kiểm tra version thực tế đã cài**: `pnpm ls tailwindcss` — phải ra
   `3.4.x`, không phải `4.x`.
4. Đảm bảo `app/globals.css` được import đúng một lần trong
   `app/layout.tsx` (đã có sẵn) và không có file `globals.css` thứ hai nào
   đè lên.

## 5. Kết nối backend

Toàn bộ trang hiện đang dùng dữ liệu mock trực tiếp trong Server Component
(đánh dấu `// TODO: thay bằng gọi API thật`). Khi backend sẵn sàng:

1. Cập nhật `NEXT_PUBLIC_API_URL` trong `.env.local`.
2. Thay các hàm `getCourses()`, `getCourse()`, `getLessonMDX()`,
   `getExercise()` trong từng `page.tsx` bằng lời gọi tương ứng trong
   `lib/api/*`.
3. Thêm xác thực (NextAuth hoặc cơ chế JWT riêng) — hiện `apiClient` đã hỗ
   trợ truyền `token`, chỉ cần lấy session ở Server Component và truyền vào.

## 6. Kế hoạch triển khai tiếp theo (Roadmap)

### Giai đoạn 0 — Hoàn thiện scaffold (0.5–1 ngày)
- [ ] `npm install`, chạy thử `npm run dev`, sửa lỗi type/import phát sinh
      do phiên bản package thực tế khác với lúc scaffold.
- [ ] Cài shadcn/ui CLI thật (`npx shadcn@latest init`) để sinh đúng file
      `components/ui/*` đầy đủ (hiện chỉ có `button`, `card`, `progress`
      viết tay tối thiểu để chạy được).
- [ ] Thiết lập CI cơ bản: lint + typecheck + unit test trên mỗi PR.

### Giai đoạn 1 — Xác thực & layout (2–3 ngày)
- [ ] Tích hợp NextAuth (hoặc auth provider của backend), hoàn thiện
      `/login`, `/register`, middleware bảo vệ route `(student)` và `admin`.
- [ ] Hoàn thiện `Navbar` (avatar, dropdown user, toggle theme) và
      `MobileNav` (đồng bộ menu với `Sidebar`).
- [ ] Trang `/profile`, `/settings` (form cập nhật thông tin, đổi mật khẩu).

### Giai đoạn 2 — Khoá học & bài học (3–5 ngày)
- [ ] Kết nối `coursesApi` với backend thật; render danh sách khoá học,
      chi tiết khoá học theo `chapters`.
- [ ] Route loader cho MDX theo `courseSlug/lessonSlug` (thay import tĩnh
      bằng cơ chế đọc `content/` linh hoạt — cân nhắc `next-mdx-remote` nếu
      nội dung không cố định tại build time).
- [ ] Điều hướng bài học trước/sau (`prevLessonSlug` / `nextLessonSlug`),
      đánh dấu hoàn thành bài học.
- [ ] Thêm `loading.tsx` / `error.tsx` / empty state cho các route course.

### Giai đoạn 3 — Bài tập & Quiz (5–7 ngày)
- [ ] Tích hợp Pyodide thật cho nút **Run** (chạy Python ngay trong browser).
- [ ] Hoàn thiện luồng **Submit**: gọi API → polling hoặc WebSocket để lấy
      kết quả từ queue/code runner → hiển thị trạng thái `pending/running/
      passed/failed`.
- [ ] Trang `/quizzes/[quizId]`: hiển thị câu hỏi, chấm điểm, cộng XP.
- [ ] Component hiển thị test case ẩn/hiện đúng theo `visibleTestCases` —
      không bao giờ render hidden test case nhận từ backend.

### Giai đoạn 4 — Gamification & Dashboard (3–4 ngày)
- [ ] Hoàn thiện `/dashboard` với dữ liệu thật: continue learning, XP,
      level, streak, recent activity, leaderboard preview.
- [ ] `/progress`: biểu đồ/list tiến độ theo từng khoá học.
- [ ] `/achievements`: danh sách huy hiệu, trạng thái mở khoá.
- [ ] `/leaderboard`: bảng xếp hạng (phân trang hoặc infinite scroll).

### Giai đoạn 5 — Khu vực Admin (5–7 ngày)
- [ ] CRUD `courses`, `lessons`, `exercises`, `test-cases`, `quizzes` (bảng +
      form, dùng React Hook Form + Zod).
- [ ] Trang `submissions`: xem/lọc bài nộp của học viên.
- [ ] Trang `users`: quản lý người dùng, phân quyền.
- [ ] Trang `analytics`: biểu đồ thống kê (số liệu hoàn thành, XP trung
      bình, v.v. — cân nhắc thư viện chart nhẹ như Recharts).
- [ ] Bảo vệ toàn bộ `/admin` bằng kiểm tra `role === 'admin'` ở middleware.

### Giai đoạn 6 — Chất lượng & Performance (song song, liên tục)
- [ ] Viết thêm unit test cho từng component UI/domain quan trọng.
- [ ] Viết đầy đủ kịch bản E2E theo luồng chuẩn: `Register → Login →
      Dashboard → Course → Lesson → Exercise → Submit → Result → XP →
      Progress` (đã có khung trong `tests/e2e/critical-flow.spec.ts`, đang
      `test.skip`).
- [ ] Audit accessibility (keyboard nav, focus state, contrast) cho các
      component tương tác (Dialog, Tabs, Dropdown).
- [ ] Kiểm tra bundle size, đảm bảo Monaco/Pyodide không chặn First Load JS
      của các trang không liên quan.
- [ ] Thêm `next/image` cho ảnh khoá học, `next/font` cho typography.

## 7. Ghi chú kỹ thuật

- Toàn bộ import dùng alias `@/*` (map tới thư mục gốc), khai báo trong
  `tsconfig.json`.
- `components.json` đã cấu hình sẵn cho `shadcn/ui` CLI — chạy
  `npx shadcn@latest add <component>` để thêm component mới đúng chuẩn dự
  án thay vì tự viết tay.
- File `types/mdx.d.ts` khai báo type cho import `.mdx` trực tiếp — cần
  thiết vì `LessonPage` hiện `import()` file MDX tĩnh (tạm thời, xem
  Giai đoạn 2).
- Không có Redux/state management toàn cục — theo đúng chuẩn skill, chỉ
  thêm khi có yêu cầu cụ thể (ví dụ giỏ hàng, notification real-time).

## 8. Liên kết tới các quy tắc gốc

Chi tiết đầy đủ về kiến trúc, coding rules, accessibility, testing nằm
trong skill `frontend-learning-platform`
(`/mnt/skills/plugins/frontend-learning-platform/SKILL.md`). File README
này là bản tóm tắt + kế hoạch thực thi cụ thể cho repo.
