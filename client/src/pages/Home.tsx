import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import {
  ArrowRight,
  Cpu,
  Github,
  Instagram,
  Linkedin,
  Terminal,
  Twitter,
  Zap,
} from "lucide-react";
import { useCallback, useEffect, useMemo, useState } from "react";
import { fetcher } from "@/lib/api";
import { apiPost, setAuthToken, getAuthToken } from "@/lib/api";
import { Article, Category, Lab, Music, PaginatedArticleResponse, RawLab, Tag } from "@/types/api";
import { Link } from "wouter";
import { GITHUB_OAUTH_URL, GOOGLE_OAUTH_URL, GITHUB_PROFILE } from "@/const";

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

const formatDate = (value?: string) => {
  if (!value) return "未发布";
  const date = new Date(value);
  return Number.isNaN(date.getTime()) ? "未发布" : date.toLocaleDateString("zh-CN");
};

export default function Home() {
  const [scrollY, setScrollY] = useState(0);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [articles, setArticles] = useState<Article[]>([]);
  const [labs, setLabs] = useState<Lab[]>([]);
  const [categories, setCategories] = useState<Category[]>([]);
  const [tags, setTags] = useState<Tag[]>([]);
  const [musics, setMusics] = useState<Music[]>([]);
  const [currentTrack, setCurrentTrack] = useState<Music | null>(null);
  const [loginMessage, setLoginMessage] = useState("");
  const [token, setToken] = useState(getAuthToken());
  const [loginForm, setLoginForm] = useState({ username: "", password: "" });

  const apiGetMusic = async () => {
    const res = await fetcher<{ musics: Music[] }>("/music");
    return res?.musics || [];
  };

  const loadData = useCallback(async () => {
    setLoading(true);
    setError(null);
    try {
      const [articleRes, labRes, categoryRes, tagRes, musicRes] = await Promise.all([
        fetcher<PaginatedArticleResponse>("/articles", { page_size: 6 }),
        fetcher<RawLab[]>("/labs"),
        fetcher<Category[]>("/categories"),
        fetcher<Tag[]>("/tags"),
        apiGetMusic(),
      ]);

      setArticles(articleRes?.articles || []);
      setLabs((labRes || []).map(normalizeLab));
      setCategories(categoryRes || []);
      setTags(tagRes || []);
      setMusics(musicRes || []);
      if ((musicRes || []).length) {
        setCurrentTrack(musicRes![0]);
      }
    } catch (err) {
      setError("获取后端数据失败，请确认 Gin 接口正常启动（默认端口 8080）。");
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    const handleScroll = () => setScrollY(window.scrollY);
    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, []);

  useEffect(() => {
    loadData();
  }, [loadData]);

  const featuredArticle = useMemo(
    () => articles.find((item) => item.is_top) ?? articles[0],
    [articles]
  );

  const latestArticles = useMemo(() => articles.slice(0, 3), [articles]);

  const totals = useMemo(
    () => ({
      views: articles.reduce((sum, item) => sum + (item.views || 0), 0),
      likes: articles.reduce((sum, item) => sum + (item.likes || 0), 0),
    }),
    [articles]
  );

  const handleLogin = async () => {
    try {
      const res = await apiPost<{ token?: string }>("/auth/login", {
        username: loginForm.username,
        password: loginForm.password,
      });
      if ((res as any).token) {
        setAuthToken((res as any).token);
        setToken((res as any).token);
        setLoginMessage("登录成功，Token 已保存");
      } else {
        setLoginMessage("登录成功，但未返回 Token");
      }
    } catch (err) {
      setLoginMessage("登录失败，请检查账号密码");
    }
  };

  return (
    <div className="min-h-screen overflow-x-hidden bg-background text-foreground selection:bg-primary selection:text-black font-mono">
      <div className="fixed inset-0 -z-10 overflow-hidden pointer-events-none">
        <div
          className="absolute top-0 left-0 w-full h-full opacity-60"
          style={{
            backgroundImage: "url(/images/cyber-bg.png)",
            backgroundSize: "cover",
            backgroundPosition: "center",
            transform: `translateY(${scrollY * 0.1}px) scale(${1 + scrollY * 0.0002})`,
            filter: "contrast(1.2) brightness(0.8)",
          }}
        />
        <div className="absolute inset-0 bg-background/80 mix-blend-multiply" />
        <div className="absolute inset-0 bg-[url('/images/cyber-shape-2.png')] opacity-10 mix-blend-screen animate-pulse" />
      </div>

      <div className="fixed inset-0 z-40 pointer-events-none">
        <div className="absolute top-32 left-6 md:left-12 flex flex-col gap-2 opacity-70">
          <div className="flex items-center gap-2 text-[10px] text-primary font-mono tracking-widest">
            <div className="w-2 h-2 bg-primary animate-ping" />
            API STATUS · {loading ? "同步中..." : "已连接"}
          </div>
          <div className="h-px w-40 bg-gradient-to-r from-primary to-transparent" />
          <div className="text-[10px] text-muted-foreground font-mono">
            ARTICLES: {articles.length} · LABS: {labs.length} · TAGS: {tags.length}
          </div>
        </div>

        <div className="absolute top-32 right-6 md:right-12 flex flex-col items-end gap-2 opacity-70">
          <div className="flex items-center gap-1">
            {[articles.length, categories.length, tags.length, labs.length, totals.views].map(
              (value, i) => (
                <div
                  key={i}
                  className="w-1 h-4 bg-secondary/50"
                  style={{ opacity: Math.min(1, 0.25 + Math.min(value, 12) / 15) }}
                />
              )
            )}
          </div>
          <div className="text-[10px] text-secondary font-mono tracking-widest">
            端口: 8080 // 数据最新: {!loading && !error ? "是" : "否"}
          </div>
        </div>

        <div className="absolute top-0 left-8 md:left-16 w-px h-screen bg-white/5" />
        <div className="absolute top-0 right-8 md:right-16 w-px h-screen bg-white/5" />
      </div>

      <nav className="fixed top-0 left-0 right-0 z-50 px-6 py-6 border-b border-primary/20 bg-background/80 backdrop-blur-md">
        <div className="container mx-auto flex items-center justify-between">
          <a
            href="/"
            className="text-2xl font-bold tracking-tighter hover:text-primary transition-colors flex items-center gap-2"
          >
            <Terminal className="w-6 h-6 text-primary" />
            <span className="glitch" data-text="Gin Blog">
              Gin Blog
            </span>
          </a>
          <div className="hidden md:flex items-center gap-8">
            {[
              { label: "专题", href: "#work" },
              { label: "概览", href: "#about" },
              { label: "文章", href: "#blog" },
              { label: "接口", href: "#contact" },
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
          <div className="flex items-center gap-3">
            <Link href="/admin">
              <Button variant="ghost" className="rounded-none border border-primary/40 text-primary hover:bg-primary hover:text-black">
                管理面板
              </Button>
            </Link>
            <Link href="/publish">
              <Button variant="ghost" className="rounded-none border border-secondary/40 text-secondary hover:bg-secondary hover:text-black">
                发布文章
              </Button>
            </Link>
            <Button
              variant="outline"
              className="btn-neon rounded-none border-primary text-primary hover:bg-primary hover:text-black"
              onClick={loadData}
            >
              重新拉取数据
            </Button>
          </div>
        </div>
      </nav>

      <section className="relative min-h-screen flex items-center pt-20 overflow-hidden">
        <div className="container mx-auto grid lg:grid-cols-2 gap-12 items-center relative z-10">
          <div className="space-y-8">
            <div className="inline-flex items-center gap-2 px-4 py-1 border border-secondary/50 bg-secondary/10 text-xs font-bold text-secondary uppercase tracking-[0.2em]">
              <span className="w-2 h-2 bg-secondary animate-pulse" />
              {featuredArticle?.category?.name || "博客系统"} · {loading ? "同步中" : "实时数据"}
            </div>

            <h1 className="text-5xl md:text-7xl font-bold leading-[0.9] tracking-tighter">
              <span className="block text-white mb-2">
                {featuredArticle?.title || "连接 Gin 博客后端"}
              </span>
              <span className="glitch text-primary block" data-text="后端数据直连">
                后端数据直连
              </span>
            </h1>

            <p className="text-lg md:text-xl text-muted-foreground max-w-lg leading-relaxed border-l-2 border-primary/50 pl-6">
              {featuredArticle?.excerpt ||
                "使用 React 保留现有布局，直接消费 Gin API，文章、分类、标签数据均由数据库驱动。"}
            </p>

            <div className="flex flex-wrap gap-4 pt-4">
              <a href="#work">
                <Button size="lg" className="btn-neon h-14 px-8 text-lg rounded-none">
                  查看实验室 <Cpu className="ml-2 h-5 w-5" />
                </Button>
              </a>
              <a href="#blog">
                <Button
                  size="lg"
                  variant="outline"
                  className="btn-neon-cyan h-14 px-8 text-lg rounded-none border-secondary text-secondary hover:bg-secondary hover:text-black"
                >
                  最新文章
                </Button>
              </a>
              <Button
                size="lg"
                variant="outline"
                className="btn-neon h-14 px-8 text-lg rounded-none border-primary text-primary hover:bg-primary hover:text-black"
                onClick={handleLogin}
              >
                本地登录
              </Button>
            </div>

            <div className="flex items-center gap-6 pt-4 text-muted-foreground">
              <div className="text-xs uppercase tracking-widest">
                阅读 {totals.views} · 点赞 {totals.likes}
              </div>
            </div>

            <div className="flex items-center gap-6 pt-4 text-muted-foreground">
              <a
                href={GITHUB_PROFILE}
                target="_blank"
                rel="noreferrer"
                className="hover:text-primary hover:scale-110 transition-all duration-300 p-2 border border-transparent hover:border-primary/50 hover:bg-primary/10"
              >
                <Github className="h-5 w-5" />
              </a>
              <a
                href={GITHUB_OAUTH_URL}
                target="_blank"
                rel="noreferrer"
                className="hover:text-primary hover:scale-110 transition-all duration-300 p-2 border border-transparent hover:border-primary/50 hover:bg-primary/10"
              >
                GitHub 登录
              </a>
              <a
                href={GOOGLE_OAUTH_URL}
                target="_blank"
                rel="noreferrer"
                className="hover:text-secondary hover:scale-110 transition-all duration-300 p-2 border border-transparent hover:border-secondary/50 hover:bg-secondary/10"
              >
                Google 登录
              </a>
            </div>

            {loginMessage && <div className="text-sm text-primary pt-2">{loginMessage}</div>}
          </div>

          <div className="relative hidden lg:flex h-[600px] w-full items-center justify-center">
            <img
              src="/images/cyber-shape-1.png"
              alt="数据图形"
              className="max-w-full max-h-[90%] object-contain animate-float drop-shadow-[0_0_50px_rgba(255,0,255,0.3)] z-20"
              style={{ animationDuration: "4s" }}
            />
            <div className="absolute inset-0 bg-gradient-to-t from-background via-transparent to-transparent z-10 pointer-events-none" />
          </div>
        </div>

        <div className="absolute bottom-0 left-0 w-full h-px bg-gradient-to-r from-transparent via-primary to-transparent opacity-50" />
        <div className="absolute top-1/4 right-0 w-64 h-64 bg-primary/20 blur-[100px] rounded-full pointer-events-none" />
        <div className="absolute bottom-1/4 left-0 w-64 h-64 bg-secondary/20 blur-[100px] rounded-full pointer-events-none" />
      </section>

      <section id="work" className="py-32 relative border-t border-white/5">
        <div className="container mx-auto">
          <div className="flex flex-col md:flex-row justify-between items-end mb-16 gap-6">
            <div>
              <h2 className="text-4xl md:text-5xl font-bold mb-4 flex items-center gap-4">
                <span className="text-primary text-2xl">01.</span> 数据实验室
              </h2>
              <p className="text-muted-foreground text-lg max-w-md font-mono">
                来自 labs 表的专题数据，保留原有栅格结构，只替换为后端内容。
              </p>
            </div>
            <Button
              variant="ghost"
              className="group text-lg hover:text-primary hover:bg-transparent"
              onClick={loadData}
            >
              刷新专题
              <ArrowRight className="ml-2 h-5 w-5 transition-transform group-hover:translate-x-1" />
            </Button>
          </div>

          <div className="grid md:grid-cols-2 gap-8">
            {labs.slice(0, 4).map((item, index) => (
              <Link
                href={`/labs/${item.slug}`}
                key={item.id}
                className={`group relative ${index % 2 === 1 ? "md:translate-y-24" : ""}`}
              >
                <Card className="cyber-panel rounded-none border-0 bg-black/50 overflow-visible">
                  <CardContent className="p-0 relative aspect-[4/3] overflow-hidden">
                    <img
                      src={item.hero_image || `/images/project-${(index % 4) + 1}.png`}
                      alt={item.title}
                      className="w-full h-full object-cover transition-transform duration-700 group-hover:scale-110 group-hover:filter group-hover:contrast-125"
                    />
                    <div className="absolute inset-0 bg-gradient-to-t from-background via-background/20 to-transparent opacity-60" />

                    <div className="absolute inset-0 bg-[url('/images/cyber-shape-2.png')] opacity-0 group-hover:opacity-20 transition-opacity duration-100 mix-blend-overlay" />

                    <div className="absolute bottom-0 left-0 right-0 p-6 bg-black/90 border-t border-primary/30 backdrop-blur-md translate-y-0 transition-transform duration-300">
                      <div className="flex justify-between items-center mb-2">
                        <div className="text-xs text-primary font-mono">
                          LAB_{String(item.id).padStart(2, "0")} // {item.badge || "LIVE"}
                        </div>
                        <div className="h-1 w-12 bg-secondary/30 overflow-hidden">
                          <div className="h-full w-1/2 bg-secondary animate-pulse" />
                        </div>
                      </div>
                      <h3 className="text-xl font-bold mb-2 group-hover:text-primary transition-colors truncate">
                        {item.title}
                      </h3>
                      <p className="text-muted-foreground text-sm mb-4 line-clamp-2">
                        {item.subtitle || item.description || "后端暂无描述"}
                      </p>
                      <div className="flex gap-2 flex-wrap">
                        {item.badge && (
                          <span className="text-[10px] px-2 py-1 border border-white/10 text-white/60 bg-white/5">
                            {item.badge}
                          </span>
                        )}
                        {item.focus && (
                          <span className="text-[10px] px-2 py-1 border border-white/10 text-white/60 bg-white/5">
                            {item.focus}
                          </span>
                        )}
                      </div>
                    </div>

                    <div className="absolute top-0 left-0 w-4 h-4 border-t-2 border-l-2 border-primary transition-all duration-300 group-hover:w-8 group-hover:h-8" />
                    <div className="absolute top-0 right-0 w-4 h-4 border-t-2 border-r-2 border-primary transition-all duration-300 group-hover:w-8 group-hover:h-8" />
                    <div className="absolute bottom-0 left-0 w-4 h-4 border-b-2 border-l-2 border-primary transition-all duration-300 group-hover:w-8 group-hover:h-8" />
                    <div className="absolute bottom-0 right-0 w-4 h-4 border-b-2 border-r-2 border-primary transition-all duration-300 group-hover:w-8 group-hover:h-8" />
                  </CardContent>
                </Card>
              </Link>
            ))}

            {!labs.length && !loading && (
              <div className="col-span-2 text-center text-muted-foreground">
                暂无实验室数据，请在后端填充 labs 表。
              </div>
            )}
          </div>
        </div>
      </section>

      <section id="about" className="py-32 relative overflow-hidden bg-black/50">
        <div className="absolute inset-0 bg-[url('/images/cyber-shape-2.png')] opacity-5" />

        <div className="container mx-auto relative z-10">
          <div className="max-w-4xl mx-auto text-center">
            <span className="text-secondary font-bold tracking-[0.3em] uppercase mb-6 block">数据概览</span>
            <h2 className="text-3xl md:text-5xl font-bold leading-tight mb-12">
              React 前端 + Gin 后端，实时读取数据库表内容
            </h2>

            <div className="grid md:grid-cols-2 gap-12 text-lg text-muted-foreground leading-relaxed text-left border border-white/10 p-8 bg-black/40 backdrop-blur-sm relative">
              <div className="absolute -top-1 -left-1 w-3 h-3 bg-primary" />
              <div className="absolute -bottom-1 -right-1 w-3 h-3 bg-secondary" />

              <p>
                <span className="text-primary font-bold">&gt;</span> 页面结构保持不变，内容来自
                /api/articles、/api/categories、/api/tags、/api/labs 等接口。
              </p>
              <p>
                <span className="text-secondary font-bold">&gt;</span> 支持 MySQL/SQLite，前端通过 Axios 直接消费 JSON，
                兼容已有主题样式与交互。
              </p>
            </div>

            <div className="mt-16 grid grid-cols-2 md:grid-cols-4 gap-4">
              {[
                { label: "已发布文章", value: articles.length },
                { label: "分类数量", value: categories.length },
                { label: "标签数量", value: tags.length },
                { label: "累计阅读", value: totals.views },
              ].map((stat, i) => (
                <div
                  key={i}
                  className="border border-white/10 p-6 hover:border-primary/50 hover:bg-primary/5 transition-all group"
                >
                  <div className="text-3xl font-bold text-white mb-2 group-hover:text-primary group-hover:scale-110 transition-all duration-300">
                    {stat.value}
                  </div>
                  <div className="text-xs text-muted-foreground uppercase tracking-widest">{stat.label}</div>
                </div>
              ))}
            </div>
          </div>
        </div>
      </section>

      <section id="blog" className="py-32 relative">
        <div className="container mx-auto">
          <h2 className="text-4xl md:text-5xl font-bold mb-16 text-center flex items-center justify-center gap-4">
            <span className="text-secondary text-2xl">02.</span> 最新文章
          </h2>

          <div className="grid md:grid-cols-3 gap-8">
            {latestArticles.map((post) => (
              <Link key={post.id} href={`/articles/${post.id}`} className="group block">
                <article className="cyber-panel h-full p-6 hover:border-secondary/50 transition-all duration-300 hover:-translate-y-2">
                  <div className="mb-6 aspect-video bg-black border border-white/10 overflow-hidden relative">
                    <div className="absolute inset-0 bg-gradient-to-br from-primary/20 to-secondary/20 group-hover:opacity-100 opacity-50 transition-opacity" />
                    <Zap className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 text-white/20 w-12 h-12 group-hover:text-secondary group-hover:scale-125 transition-all duration-500" />
                  </div>
                  <div className="flex items-center gap-4 text-xs font-bold uppercase tracking-wider text-muted-foreground mb-4">
                    <span className="text-secondary px-2 py-1 border border-secondary/30 bg-secondary/10">
                      {post.tags?.[0]?.name || post.category?.name || "未分类"}
                    </span>
                    <span className="font-mono">{formatDate(post.published_at || post.created_at)}</span>
                  </div>
                  <h3 className="text-xl font-bold mb-4 group-hover:text-primary transition-colors">
                    {post.title}
                  </h3>
                  <p className="text-sm text-muted-foreground font-mono line-clamp-3">
                    {post.excerpt || "这篇文章来自数据库，可在后端编辑或新增。"}
                  </p>
                </article>
              </Link>
            ))}

            {!latestArticles.length && !loading && (
              <div className="col-span-3 text-center text-muted-foreground">
                暂无文章数据，请通过 /api/articles 创建内容。
              </div>
            )}
          </div>
        </div>
      </section>

      {/* Music player */}
      <section className="py-20 border-t border-primary/20 bg-black/50">
        <div className="container mx-auto">
          <div className="flex items-center justify-between mb-8">
            <h2 className="text-3xl md:text-4xl font-bold flex items-center gap-3">
              <span className="text-secondary text-2xl">03.</span> 音乐播放
            </h2>
            <Button variant="ghost" onClick={loadData}>
              刷新曲目
            </Button>
          </div>
          <div className="grid md:grid-cols-3 gap-6">
            <div className="md:col-span-2 space-y-4">
              <div className="border border-white/10 p-4 bg-black/60">
                <div className="flex items-center justify-between mb-3">
                  <div>
                    <div className="text-lg font-bold">{currentTrack?.title || "未选择曲目"}</div>
                    <div className="text-xs text-muted-foreground">{currentTrack?.artist}</div>
                  </div>
                  <span className="text-xs text-muted-foreground">
                    {currentTrack ? `${currentTrack.duration || 0}s` : "--"}
                  </span>
                </div>
                <audio
                  key={currentTrack?.id || "audio"}
                  controls
                  className="w-full"
                  src={currentTrack?.url}
                />
              </div>
            </div>
            <div className="space-y-3">
              {(musics || []).map((m) => (
                <button
                  key={m.id}
                  className={`w-full text-left border p-3 text-sm hover:border-primary transition-colors ${
                    currentTrack?.id === m.id ? "border-primary" : "border-white/10"
                  }`}
                  onClick={() => setCurrentTrack(m)}
                >
                  <div className="font-bold">{m.title}</div>
                  <div className="text-xs text-muted-foreground">{m.artist}</div>
                </button>
              ))}
              {!musics.length && <div className="text-muted-foreground text-sm">暂无音乐数据</div>}
            </div>
          </div>
        </div>
      </section>

      <footer className="py-20 border-t border-primary/20 bg-black relative overflow-hidden" id="contact">
        <div className="container mx-auto text-center relative z-10">
          {error ? (
            <div className="text-red-400 mb-6">{error}</div>
          ) : (
            <div className="text-muted-foreground mb-6">
              数据来源：Gin 后端（端口 8080），前端通过 Axios 直接访问。
            </div>
          )}

          <h2 className="text-4xl md:text-6xl font-bold mb-8">
            保持 <span className="text-primary">同源数据</span> 连接
          </h2>
          <Button size="lg" className="btn-neon h-16 px-12 text-xl rounded-none mb-16" onClick={loadData}>
            立即同步
          </Button>

          <div className="flex justify-between items-end pt-16 border-t border-white/10">
            <div className="text-left">
              <div className="text-2xl font-bold mb-2 flex items-center gap-2">
                <Terminal className="w-5 h-5 text-primary" /> Gin Blog
              </div>
              <p className="text-muted-foreground text-xs font-mono">
                © 2024 Gin Blog. React 前端已切换完成，Vue 模块已移除。
              </p>
            </div>
            <div className="flex gap-6 text-xs font-mono uppercase tracking-widest">
              <a href="#" className="text-muted-foreground hover:text-primary transition-colors">
                [ API /api/articles ]
              </a>
              <a href="#" className="text-muted-foreground hover:text-primary transition-colors">
                [ API /api/labs ]
              </a>
            </div>
          </div>
        </div>
      </footer>
    </div>
  );
}
