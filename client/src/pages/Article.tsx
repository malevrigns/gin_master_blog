import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { fetcher } from "@/lib/api";
import { Article } from "@/types/api";
import { ArrowLeft, Loader2, Terminal, Zap } from "lucide-react";
import { useEffect, useMemo, useState } from "react";
import { Link, useParams } from "wouter";

const formatDate = (value?: string) => {
  if (!value) return "未发布";
  const date = new Date(value);
  return Number.isNaN(date.getTime()) ? "未发布" : date.toLocaleDateString("zh-CN");
};

export default function ArticlePage() {
  const params = useParams<{ id: string }>();
  const [article, setArticle] = useState<Article | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const load = async () => {
      setLoading(true);
      setError(null);
      try {
        const data = await fetcher<Article>(`/articles/${params.id}`);
        setArticle(data);
      } catch (err) {
        setError("未找到该文章，或后端接口不可用。");
      } finally {
        setLoading(false);
      }
    };
    load();
  }, [params.id]);

  const tags = useMemo(() => article?.tags || [], [article]);

  return (
    <div className="min-h-screen overflow-x-hidden bg-background text-foreground selection:bg-primary selection:text-black font-mono">
      {/* 背景与装饰层 */}
      <div className="fixed inset-0 -z-10 overflow-hidden pointer-events-none">
        <div
          className="absolute top-0 left-0 w-full h-full opacity-60"
          style={{
            backgroundImage: "url(/images/cyber-bg.png)",
            backgroundSize: "cover",
            backgroundPosition: "center",
            filter: "contrast(1.2) brightness(0.8)",
          }}
        />
        <div className="absolute inset-0 bg-background/80 mix-blend-multiply" />
        <div className="absolute inset-0 bg-[url('/images/cyber-shape-2.png')] opacity-10 mix-blend-screen animate-pulse" />
      </div>

      {/* HUD 层，与首页保持一致的感觉 */}
      <div className="fixed inset-0 z-40 pointer-events-none">
        <div className="absolute top-32 left-6 md:left-12 flex flex-col gap-2 opacity-70">
          <div className="flex items-center gap-2 text-[10px] text-primary font-mono tracking-widest">
            <div className="w-2 h-2 bg-primary animate-ping" />
            ARTICLE VIEW · {loading ? "同步中..." : article ? "已连接" : "未获取"}
          </div>
          <div className="h-px w-40 bg-gradient-to-r from-primary to-transparent" />
          <div className="text-[10px] text-muted-foreground font-mono">
            ID: {params.id} · {article?.category?.name || "未分类"} ·{" "}
            {article ? `TAGS: ${tags.length}` : "加载中"}
          </div>
        </div>
        <div className="absolute top-32 right-6 md:right-12 flex flex-col items-end gap-2 opacity-70">
          <div className="flex items-center gap-1">
            {[article?.views || 0, article?.likes || 0, tags.length || 1, 5, 8].map((value, i) => (
              <div
                key={i}
                className="w-1 h-4 bg-secondary/50"
                style={{ opacity: Math.min(1, 0.25 + Math.min(value, 12) / 15) }}
              />
            ))}
          </div>
          <div className="text-[10px] text-secondary font-mono tracking-widest">
            {article ? formatDate(article.published_at || article.created_at) : "等待数据"}
          </div>
        </div>
        <div className="absolute top-0 left-8 md:left-16 w-px h-screen bg-white/5" />
        <div className="absolute top-0 right-8 md:right-16 w-px h-screen bg-white/5" />
      </div>

      {/* 导航保持首页风格 */}
      <nav className="fixed top-0 left-0 right-0 z-50 px-6 py-6 border-b border-primary/20 bg-background/80 backdrop-blur-md">
        <div className="container mx-auto flex items-center justify-between">
          <Link href="/" className="text-2xl font-bold tracking-tighter hover:text-primary transition-colors flex items-center gap-2">
            <Terminal className="w-6 h-6 text-primary" />
            <span className="glitch" data-text="Gin Blog">
              Gin Blog
            </span>
          </Link>
          <div className="hidden md:flex items-center gap-8">
            {[
              { label: "专题", href: "/#work" },
              { label: "概览", href: "/#about" },
              { label: "文章", href: "/#blog" },
              { label: "接口", href: "/#contact" },
            ].map((item) => (
              <a
                key={item.label}
                href={item.href}
                className="text-sm font-bold uppercase tracking-widest hover:text-primary transition-colors relative group"
              >
                <span className="text-primary/50 mr-1">&lt;</span>
                {item.label}
                <span className="text-primary/50 ml-1">/&gt;</span>
              </a>
            ))}
          </div>
          <Link href="/">
            <Button variant="outline" className="btn-neon rounded-none border-primary text-primary hover:bg-primary hover:text-black">
              返回首页
            </Button>
          </Link>
        </div>
      </nav>

      {/* Hero 区域 */}
      <section className="relative min-h-screen flex items-center pt-24 md:pt-28">
        <div className="container mx-auto grid lg:grid-cols-2 gap-12 items-center relative z-10 px-4">
          <div className="space-y-8">
            <div className="inline-flex items-center gap-2 px-4 py-1 border border-secondary/50 bg-secondary/10 text-xs font-bold text-secondary uppercase tracking-[0.2em]">
              <span className="w-2 h-2 bg-secondary animate-pulse" />
              {article?.category?.name || "博客文章"} · {loading ? "同步中" : "实时数据"}
            </div>
            <h1 className="text-4xl md:text-6xl font-bold leading-[0.95] tracking-tighter">
              <span className="block text-white mb-3">{article?.title || "加载中..."}</span>
              <span className="glitch text-primary block" data-text="ARTICLE DETAIL">
                ARTICLE DETAIL
              </span>
            </h1>
            <p className="text-lg md:text-xl text-muted-foreground leading-relaxed border-l-2 border-primary/50 pl-6">
              {article?.excerpt || "来自 Gin 后端的文章详情，保持与首页一致的赛博视觉风格。"}
            </p>
            <div className="flex flex-wrap gap-3">
              {(tags.length ? tags : [{ id: -1, name: "未分类", slug: "none" }]).map((tag) => (
                <span
                  key={tag.id}
                  className="text-[10px] px-3 py-1 border border-white/10 text-white/80 bg-white/5 uppercase tracking-widest"
                >
                  {tag.name}
                </span>
              ))}
            </div>
            <div className="flex items-center gap-6 text-sm text-muted-foreground">
              <span>发布时间：{article ? formatDate(article.published_at || article.created_at) : "..."}</span>
              <span>阅读 {article?.views ?? 0}</span>
              <span>点赞 {article?.likes ?? 0}</span>
            </div>
            <div className="flex gap-4">
              <Link href="/">
                <Button size="lg" className="btn-neon h-12 px-6 rounded-none">回到首页</Button>
              </Link>
              <a href="/#blog">
                <Button size="lg" variant="outline" className="btn-neon-cyan h-12 px-6 rounded-none border-secondary text-secondary hover:bg-secondary hover:text-black">
                  返回文章列表
                </Button>
              </a>
            </div>
          </div>

          <div className="relative hidden lg:flex h-[520px] w-full items-center justify-center">
            <div className="absolute inset-0 bg-gradient-to-br from-primary/20 via-transparent to-secondary/30 blur-[100px] rounded-full" />
            <div className="relative w-full h-full border border-primary/30 bg-black/60 flex items-center justify-center overflow-hidden">
              {article?.cover_image ? (
                <img
                  src={article.cover_image}
                  alt={article.title}
                  className="w-full h-full object-cover"
                />
              ) : (
                <Zap className="w-16 h-16 text-primary/60" />
              )}
              <div className="absolute inset-0 bg-gradient-to-t from-background via-background/20 to-transparent" />
            </div>
          </div>
        </div>

        <div className="absolute bottom-0 left-0 w-full h-px bg-gradient-to-r from-transparent via-primary to-transparent opacity-50" />
        <div className="absolute top-1/4 right-0 w-64 h-64 bg-primary/20 blur-[100px] rounded-full pointer-events-none" />
        <div className="absolute bottom-1/4 left-0 w-64 h-64 bg-secondary/20 blur-[100px] rounded-full pointer-events-none" />
      </section>

      {/* 正文内容区域 */}
      <section className="py-24 relative z-10">
        <div className="container mx-auto px-4">
          {loading && (
            <div className="flex items-center gap-3 text-muted-foreground">
              <Loader2 className="h-5 w-5 animate-spin" />
              正在加载文章...
            </div>
          )}
          {error && <div className="text-red-400">{error}</div>}
          {!loading && article && (
            <Card className="cyber-panel rounded-none border-0 bg-black/60 overflow-hidden">
              <CardContent className="p-6 md:p-10 space-y-6">
                <div className="flex flex-wrap gap-3 text-xs uppercase tracking-widest text-muted-foreground">
                  <span className="px-2 py-1 border border-primary/40 text-primary bg-primary/10">
                    {article.category?.name || "未分类"}
                  </span>
                  {tags.map((tag) => (
                    <span key={tag.id} className="px-2 py-1 border border-white/10 text-white/70 bg-white/5">
                      {tag.name}
                    </span>
                  ))}
                  <span className="px-2 py-1 border border-secondary/40 text-secondary bg-secondary/10">
                    阅读 {article.views}
                  </span>
                  <span className="px-2 py-1 border border-secondary/40 text-secondary bg-secondary/10">
                    点赞 {article.likes}
                  </span>
                </div>

                <div className="prose prose-invert max-w-none leading-7 space-y-4">
                  <pre className="whitespace-pre-wrap text-sm bg-black/40 p-4 border border-white/5 rounded-none">
                    {article.content || "暂无内容"}
                  </pre>
                </div>
              </CardContent>
            </Card>
          )}
        </div>
      </section>
    </div>
  );
}
