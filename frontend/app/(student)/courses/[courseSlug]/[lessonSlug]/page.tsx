import { mdxComponents } from '@/components/mdx/mdx-components';

interface Props {
  params: { courseSlug: string; lessonSlug: string };
}

async function getLessonMDX(courseSlug: string, lessonSlug: string) {
  // TODO: đọc file MDX từ content/ theo courseSlug/lessonSlug (fs + next-mdx-remote,
  // hoặc dùng @next/mdx với route tĩnh nếu cấu trúc content cố định).
  const MDXContent = (await import(`@/content/python/basic/gioi-thieu.mdx`)).default;
  return MDXContent;
}

export default async function LessonPage({ params }: Props) {
  const MDXContent = await getLessonMDX(params.courseSlug, params.lessonSlug);

  return (
    <article className="container prose max-w-3xl py-10 dark:prose-invert">
      <MDXContent components={mdxComponents} />
    </article>
  );
}
