import { Callout, Tip, Warning, Note } from './callout';
import { CodeBlock, CodeExample } from './code-block';
import { ExerciseEmbed, QuizEmbed } from './embeds';
import { Steps, Definition } from './steps';

export const mdxComponents = {
  Callout,
  Tip,
  Warning,
  Note,
  CodeBlock,
  CodeExample,
  ExerciseEmbed,
  QuizEmbed,
  Steps,
  Definition,
  pre: CodeBlock,
};
