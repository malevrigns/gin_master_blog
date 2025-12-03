import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { fetcher } from "@/lib/api";
import { Article, Lab, RawLab } from "@/types/api";
import { ArrowLeft, Loader2, Zap } from "lucide-react";
import { useEffect, useMemo, useState } from "react";
import { Link, useParams } from "wouter";

function parseJsonField<T>(value: unknown): T | null {
  if (!value) return null;
  if (Array.isArray(value)) return value as T;
  if (typeof value === "string") {
    try {
      return JSON.parse(value) as T;
    } catch {
      return null;
    }
  }
  if (typeof value === "object") {
    try {
      return JSON.parse(JSON.stringify(value)) as T;
    } catch {
      return null;
    }
  }
  return null;
}

const normalizeLab = (lab: RawLab): Lab => ({
  ...lab,
  highlights: parseJsonField(lab.highlights),
  resource_links: parseJsonField(lab.resource_links),
});

export default function LabPage() {
  const { slug } = useParams<{ slug: string }>();
  const [lab, setLab] = useState<Lab | null>(null);
  const [articles, setArticles] = useState<Article[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const load = async () => {
      setLoading(true);
      setError(null);
      try {
        const [labRes, articleRes] = await Promise.all([
          fetcher<RawLab>(`/labs/${slug}`),
          fetcher<Article[]>(`/labs/${slug}/articles`),
        ]);
        setLab(normalizeLab(labRes));
        setArticles(articleRes || []);
      } catch (err) {
        setError("未找到该实验室或接口不可用。");
      } finally {
        setLoading(false);
      }
    };
    load();
  }, [slug]);

  const badgeText = useMemo(() => lab?.badge || "LAB", [lab]);

  return (
    <div className="min-h-screen bg-background text-foreground selection:bg-primary selection:text-black font-mono">
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

      <div className="container mx-auto px-4 py-8 relative z-10">
        <div className="flex items-center justify-between mb-6">
          <Link href="/">
            <Button variant="ghost" className="gap-2">
              <ArrowLeft className="h-4 w-4" />
              返回首页
            </Button>
          </Link>
          <div className="text-xs text-muted-foreground uppercase tracking-widest">
            {badgeText} · {lab?.focus || "Lab"}
          </div>
        </div>

        {loading && (
          <div className="flex items-center gap-3 text-muted-foreground">
            <Loader2 className="h-5 w-5 animate-spin" />
            正在加载实验室数据...
          </div>
        )}

        {error && <div className="text-red-400">{error}</div>}

        {!loading && lab && (
          <Card className="cyber-panel rounded-none border border-primary/30 bg-black/60 mb-10">
            <CardContent className="p-6 md:p-10 space-y-4">
              <div className="flex flex-wrap items-center gap-3 text-xs uppercase tracking-widest text-muted-foreground">
                <span className="px-2 py-1 border border-primary/40 text-primary bg-primary/10">{badgeText}</span>
                {lab.focus && <span className="px-2 py-1 border border-white/10">{lab.focus}</span>}
              </div>
              <h1 className="text-3xl md:text-5xl font-bold leading-tight">{lab.title}</h1>
              {lab.subtitle && <p className="text-lg text-muted-foreground">{lab.subtitle}</p>}
              {lab.hero_image && (
                <div className="border border-white/10 overflow-hidden">
                  <img src={lab.hero_image} alt={lab.title} className="w-full max-h-[420px] object-cover" />
                </div>
              )}
              {lab.description && <p className="text-muted-foreground leading-relaxed">{lab.description}</p>}

              {lab.highlights && lab.highlights.length > 0 && (
                <div className="grid md:grid-cols-2 gap-3">
                  {lab.highlights.map((h, idx) => (
                    <div key={idx} className="border border-white/10 p-3 bg-black/40">
                      <div className="text-sm font-bold text-primary">{h.title}</div>
                      <div className="text-xs text-muted-foreground">{h.description}</div>
                      {h.tag && <div className="text-[10px] uppercase mt-1 text-secondary">{h.tag}</div>}
                    </div>
                  ))}
                </div>
              )}

              {lab.resource_links && lab.resource_links.length > 0 && (
                <div className="space-y-2">
                  <div className="text-sm font-bold">资源链接</div>
                  {lab.resource_links.map((r, idx) => (
                    <a
                      key={idx}
                      href={r.url}
                      target="_blank"
                      rel="noreferrer"
                      className="flex items-center justify-between border border-white/10 px-3 py-2 hover:border-primary"
                    >
                      <div>
                        <div className="text-sm font-bold">{r.title}</div>
                        <div className="text-xs text-muted-foreground">{r.desc}</div>
                      </div>
                      <Zap className="h-4 w-4 text-primary" />
                    </a>
                  ))}
                </div>
              )}
            </CardContent>
          </Card>
        )}

        {!loading && !error && (
          <Card className="rounded-none border border-primary/30 bg-black/60">
            <CardContent className="p-4 space-y-4">
              <div className="flex items-center gap-2 text-primary font-bold uppercase tracking-widest">
                <Zap className="h-4 w-4" /> 相关文章
              </div>
              <div className="grid md:grid-cols-2 gap-4">
                {articles.map((a) => (
                  <Link key={a.id} href={`/articles/${a.id}`} className="border border-white/10 p-4 hover:border-primary">
                    <div className="text-sm text-secondary uppercase mb-1">
                      {a.category?.name || "未分类"}
                    </div>
                    <div className="text-lg font-bold mb-1">{a.title}</div>
                    <div className="text-xs text-muted-foreground line-clamp-2">{a.excerpt}</div>
                  </Link>
                ))}
                {!articles.length && <div className="text-muted-foreground text-sm">暂无相关文章</div>}
              </div>
            </CardContent>
          </Card>
        )}
      </div>
    </div>
  );
}
