import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import {
  apiDelete,
  apiGet,
  apiPost,
  apiPut,
  fetcher,
  getAuthToken,
  setAuthToken,
} from "@/lib/api";
import { Loader2, Terminal, Zap } from "lucide-react";
import { useCallback, useMemo, useState, type ReactNode } from "react";
import { Link } from "wouter";

type ActionHandler = () => Promise<unknown>;

const fieldClass =
  "w-full rounded-none border border-white/10 bg-black/40 px-3 py-2 text-sm text-white focus:outline-none focus:border-primary";

const sectionGap = "space-y-4";

const defaultBodies = {
  article: JSON.stringify(
    {
      title: "New Article Title",
      excerpt: "Summary",
      content: "Full content text...",
      cover_image: "",
      category_id: 1,
      tag_ids: [1],
      status: "published",
      is_top: false,
    },
    null,
    2,
  ),
  category: JSON.stringify({ name: "New Category", slug: "new-category", description: "Desc" }, null, 2),
  tag: JSON.stringify({ name: "New Tag", slug: "new-tag" }, null, 2),
  comment: JSON.stringify(
    { article_id: 1, content: "Nice post!", author: "Guest", email: "guest@example.com", website: "" },
    null,
    2,
  ),
  link: JSON.stringify(
    { name: "GitHub", url: "https://github.com", logo: "", desc: "Friend link", is_visible: true, sort: 1 },
    null,
    2,
  ),
  lab: JSON.stringify(
    {
      title: "New Lab",
      subtitle: "Subtitle",
      badge: "LAB",
      badge_color: "#34d399",
      description: "Intro text",
      focus: "Focus",
      hero_image: "",
      content: "Markdown or plain text",
    },
    null,
    2,
  ),
  music: JSON.stringify(
    {
      title: "Track title",
      artist: "Artist",
      cover: "",
      url: "https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3",
      duration: 180,
      lrc: "",
      is_public: true,
    },
    null,
    2,
  ),
};

function Panel({ title, children }: { title: string; children: ReactNode }) {
  return (
    <Card className="cyber-panel rounded-none border border-primary/30 bg-black/60">
      <CardContent className="p-4 md:p-6 space-y-4">
        <div className="flex items-center justify-between border-b border-white/10 pb-2">
          <h3 className="text-lg font-bold tracking-wide uppercase text-primary">{title}</h3>
        </div>
        {children}
      </CardContent>
    </Card>
  );
}

function tryParseJson(input: string) {
  if (!input.trim()) return {};
  try {
    return JSON.parse(input);
  } catch {
    return {};
  }
}

export default function Admin() {
  const [token, setTokenState] = useState(getAuthToken());
  const [busy, setBusy] = useState("");
  const [message, setMessage] = useState("waiting for action...");

  const [auth, setAuth] = useState({ username: "", email: "", password: "" });
  const [articleId, setArticleId] = useState("");
  const [articleBody, setArticleBody] = useState(defaultBodies.article);
  const [categoryId, setCategoryId] = useState("");
  const [categoryBody, setCategoryBody] = useState(defaultBodies.category);
  const [tagId, setTagId] = useState("");
  const [tagBody, setTagBody] = useState(defaultBodies.tag);
  const [commentId, setCommentId] = useState("");
  const [commentBody, setCommentBody] = useState(defaultBodies.comment);
  const [commentStatus, setCommentStatus] = useState("approved");
  const [linkId, setLinkId] = useState("");
  const [linkBody, setLinkBody] = useState(defaultBodies.link);
  const [labId, setLabId] = useState("");
  const [labBody, setLabBody] = useState(defaultBodies.lab);
  const [musicId, setMusicId] = useState("");
  const [musicBody, setMusicBody] = useState(defaultBodies.music);
  const [playlistId, setPlaylistId] = useState("");

  const setToken = useCallback((val: string) => {
    setAuthToken(val);
    setTokenState(val);
  }, []);

  const run = useCallback(async (label: string, handler: ActionHandler) => {
    setBusy(label);
    setMessage("...");
    try {
      const res = await handler();
      setMessage(JSON.stringify(res, null, 2));
    } catch (err) {
      if (err && typeof err === "object" && "response" in err && (err as any).response?.data) {
        setMessage(JSON.stringify((err as any).response.data, null, 2));
      } else if (err instanceof Error) {
        setMessage(err.message);
      } else {
        setMessage("unknown error");
      }
    } finally {
      setBusy("");
    }
  }, []);

  const tokenStatus = useMemo(() => (token ? "saved" : "not logged in"), [token]);

  return (
    <div className="min-h-screen overflow-x-hidden bg-background text-foreground selection:bg-primary selection:text-black font-mono">
      {/* background */}
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

      {/* HUD */}
      <div className="fixed inset-0 z-40 pointer-events-none">
        <div className="absolute top-28 left-6 md:left-12 flex flex-col gap-2 opacity-70">
          <div className="flex items-center gap-2 text-[10px] text-primary font-mono tracking-widest">
            <div className="w-2 h-2 bg-primary animate-ping" />
            ADMIN PANEL · TOKEN {tokenStatus}
          </div>
          <div className="h-px w-40 bg-gradient-to-r from-primary to-transparent" />
          <div className="text-[10px] text-muted-foreground font-mono">All backend endpoints exposed</div>
        </div>
        <div className="absolute top-28 right-6 md:right-12 flex flex-col items-end gap-2 opacity-70">
          <div className="flex items-center gap-1">
            {[token ? 12 : 4, 7, 5, 9, 11].map((v, i) => (
              <div key={i} className="w-1 h-4 bg-secondary/50" style={{ opacity: 0.3 + v / 15 }} />
            ))}
          </div>
          <div className="text-[10px] text-secondary font-mono tracking-widest">AUTH: Bearer JWT</div>
        </div>
        <div className="absolute top-0 left-8 md:left-16 w-px h-screen bg-white/5" />
        <div className="absolute top-0 right-8 md:right-16 w-px h-screen bg-white/5" />
      </div>

      {/* nav */}
      <nav className="fixed top-0 left-0 right-0 z-50 px-6 py-6 border-b border-primary/20 bg-background/80 backdrop-blur-md">
        <div className="container mx-auto flex items-center justify-between">
          <Link
            href="/"
            className="text-2xl font-bold tracking-tighter hover:text-primary transition-colors flex items-center gap-2"
          >
            <Terminal className="w-6 h-6 text-primary" />
            <span className="glitch" data-text="Gin Blog Admin">
              Gin Blog Admin
            </span>
          </Link>
          <div className="hidden md:flex items-center gap-6 text-sm uppercase tracking-widest">
            <a href="/#blog" className="hover:text-primary">
              Back to Home
            </a>
            <span className="text-primary">Publish & Manage</span>
          </div>
          <Button variant="outline" className="btn-neon rounded-none border-primary text-primary hover:bg-primary hover:text-black">
            Token: {tokenStatus}
          </Button>
        </div>
      </nav>

      <section className="pt-28 pb-16 relative">
        <div className="container mx-auto px-4 space-y-6">
          <header className="pt-10 pb-6">
            <p className="text-sm uppercase tracking-[0.4em] text-secondary mb-2">Full API Surface</p>
            <h1 className="text-4xl md:text-5xl font-bold leading-[1.1] text-white mb-3">
              Publish Articles + Operate Every Backend Endpoint
            </h1>
            <p className="text-muted-foreground max-w-3xl">
              After login you can create/update/delete articles, categories, tags, comments, links, labs, music and playlists.
              All requests carry the saved JWT token.
            </p>
          </header>

          {/* Auth */}
          <Panel title="Auth · Register / Login / Profile">
            <div className={sectionGap}>
              <div className="grid md:grid-cols-3 gap-3">
                <input
                  className={fieldClass}
                  placeholder="Username"
                  value={auth.username}
                  onChange={(e) => setAuth({ ...auth, username: e.target.value })}
                />
                <input
                  className={fieldClass}
                  placeholder="Email (for register)"
                  value={auth.email}
                  onChange={(e) => setAuth({ ...auth, email: e.target.value })}
                />
                <input
                  type="password"
                  className={fieldClass}
                  placeholder="Password"
                  value={auth.password}
                  onChange={(e) => setAuth({ ...auth, password: e.target.value })}
                />
              </div>
              <div className="flex flex-wrap gap-3">
                <Button
                  disabled={busy === "register"}
                  onClick={() =>
                    run("register", async () => {
                      const res = await apiPost<{ token?: string }>("/auth/register", auth);
                      if ((res as any).token) setToken((res as any).token);
                      return res;
                    })
                  }
                >
                  {busy === "register" && <Loader2 className="h-4 w-4 animate-spin mr-2" />}
                  Register + Save Token
                </Button>
                <Button
                  variant="secondary"
                  disabled={busy === "login"}
                  onClick={() =>
                    run("login", async () => {
                      const res = await apiPost<{ token?: string }>("/auth/login", {
                        username: auth.username,
                        password: auth.password,
                      });
                      if ((res as any).token) setToken((res as any).token);
                      return res;
                    })
                  }
                >
                  {busy === "login" && <Loader2 className="h-4 w-4 animate-spin mr-2" />}
                  Login
                </Button>
                <Button variant="outline" onClick={() => run("profile", () => apiGet("/auth/profile"))}>
                  Fetch Profile
                </Button>
                <Button variant="ghost" onClick={() => setToken("")}>
                  Clear Token
                </Button>
              </div>
            </div>
          </Panel>

          {/* Articles */}
          <Panel title="Articles · Create / Update / Delete / Like / List">
            <div className={sectionGap}>
              <div className="grid md:grid-cols-2 gap-3">
                <input
                  className={fieldClass}
                  placeholder="Article ID (for update/delete/like)"
                  value={articleId}
                  onChange={(e) => setArticleId(e.target.value)}
                />
                <textarea
                  className={`${fieldClass} h-40 font-mono`}
                  value={articleBody}
                  onChange={(e) => setArticleBody(e.target.value)}
                />
              </div>
              <div className="flex flex-wrap gap-3">
                <Button onClick={() => run("create-article", () => apiPost("/articles", tryParseJson(articleBody)))}>
                  Create
                </Button>
                <Button
                  variant="secondary"
                  onClick={() => run("update-article", () => apiPut(`/articles/${articleId}`, tryParseJson(articleBody)))}
                >
                  Update
                </Button>
                <Button variant="destructive" onClick={() => run("delete-article", () => apiDelete(`/articles/${articleId}`))}>
                  Delete
                </Button>
                <Button variant="outline" onClick={() => run("like-article", () => apiPost(`/articles/${articleId}/like`))}>
                  Like
                </Button>
                <Button variant="ghost" onClick={() => run("list-articles", () => fetcher("/articles", { page_size: 20 }))}>
                  List
                </Button>
              </div>
            </div>
          </Panel>

          {/* Categories / Tags */}
          <Panel title="Categories / Tags · CRUD">
            <div className={sectionGap}>
              <div className="grid md:grid-cols-2 gap-3">
                <input
                  className={fieldClass}
                  placeholder="Category/Tag ID (update/delete)"
                  value={categoryId}
                  onChange={(e) => {
                    setCategoryId(e.target.value);
                    setTagId(e.target.value);
                  }}
                />
                <textarea
                  className={`${fieldClass} h-32 font-mono`}
                  value={categoryBody}
                  onChange={(e) => setCategoryBody(e.target.value)}
                />
              </div>
              <div className="flex flex-wrap gap-3">
                <Button onClick={() => run("create-category", () => apiPost("/categories", tryParseJson(categoryBody)))}>
                  Create Category
                </Button>
                <Button
                  variant="secondary"
                  onClick={() => run("update-category", () => apiPut(`/categories/${categoryId}`, tryParseJson(categoryBody)))}
                >
                  Update Category
                </Button>
                <Button variant="destructive" onClick={() => run("delete-category", () => apiDelete(`/categories/${categoryId}`))}>
                  Delete Category
                </Button>
                <Button variant="outline" onClick={() => run("list-categories", () => apiGet("/categories"))}>
                  List Categories
                </Button>
              </div>
              <div className="grid md:grid-cols-2 gap-3">
                <input
                  className={fieldClass}
                  placeholder="Tag ID (delete)"
                  value={tagId}
                  onChange={(e) => setTagId(e.target.value)}
                />
                <textarea
                  className={`${fieldClass} h-24 font-mono`}
                  value={tagBody}
                  onChange={(e) => setTagBody(e.target.value)}
                />
              </div>
              <div className="flex flex-wrap gap-3">
                <Button onClick={() => run("create-tag", () => apiPost("/tags", tryParseJson(tagBody)))}>Create Tag</Button>
                <Button variant="destructive" onClick={() => run("delete-tag", () => apiDelete(`/tags/${tagId}`))}>
                  Delete Tag
                </Button>
                <Button variant="outline" onClick={() => run("list-tags", () => apiGet("/tags"))}>
                  List Tags
                </Button>
              </div>
            </div>
          </Panel>

          {/* Comments */}
          <Panel title="Comments · Create / Moderate / Delete / List">
            <div className={sectionGap}>
              <div className="grid md:grid-cols-2 gap-3">
                <input
                  className={fieldClass}
                  placeholder="Comment ID (update/delete)"
                  value={commentId}
                  onChange={(e) => setCommentId(e.target.value)}
                />
                <textarea
                  className={`${fieldClass} h-32 font-mono`}
                  value={commentBody}
                  onChange={(e) => setCommentBody(e.target.value)}
                />
              </div>
              <div className="flex flex-wrap gap-3 items-center">
                <Button onClick={() => run("create-comment", () => apiPost("/comments", tryParseJson(commentBody)))}>
                  Submit Comment
                </Button>
                <input
                  className={`${fieldClass} w-48`}
                  placeholder="Status (pending/approved/rejected)"
                  value={commentStatus}
                  onChange={(e) => setCommentStatus(e.target.value)}
                />
                <Button
                  variant="secondary"
                  onClick={() => run("update-comment", () => apiPut(`/comments/${commentId}/status`, { status: commentStatus }))}
                >
                  Update Status
                </Button>
                <Button variant="destructive" onClick={() => run("delete-comment", () => apiDelete(`/comments/${commentId}`))}>
                  Delete Comment
                </Button>
                <Button variant="outline" onClick={() => run("comments", () => apiGet("/comments", { params: { page_size: 50 } }))}>
                  List Comments
                </Button>
                <Button variant="outline" onClick={() => run("pending-comments", () => apiGet("/comments/pending"))}>
                  Pending List
                </Button>
              </div>
            </div>
          </Panel>

          {/* Links */}
          <Panel title="Links · CRUD">
            <div className={sectionGap}>
              <div className="grid md:grid-cols-2 gap-3">
                <input
                  className={fieldClass}
                  placeholder="Link ID (update/delete)"
                  value={linkId}
                  onChange={(e) => setLinkId(e.target.value)}
                />
                <textarea
                  className={`${fieldClass} h-28 font-mono`}
                  value={linkBody}
                  onChange={(e) => setLinkBody(e.target.value)}
                />
              </div>
              <div className="flex flex-wrap gap-3">
                <Button onClick={() => run("create-link", () => apiPost("/links", tryParseJson(linkBody)))}>Create</Button>
                <Button variant="secondary" onClick={() => run("update-link", () => apiPut(`/links/${linkId}`, tryParseJson(linkBody)))}>
                  Update
                </Button>
                <Button variant="destructive" onClick={() => run("delete-link", () => apiDelete(`/links/${linkId}`))}>
                  Delete
                </Button>
                <Button variant="outline" onClick={() => run("list-links", () => apiGet("/links"))}>
                  List
                </Button>
              </div>
            </div>
          </Panel>

          {/* Labs */}
          <Panel title="Labs · CRUD">
            <div className={sectionGap}>
              <div className="grid md:grid-cols-2 gap-3">
                <input
                  className={fieldClass}
                  placeholder="Lab ID (update/delete)"
                  value={labId}
                  onChange={(e) => setLabId(e.target.value)}
                />
                <textarea
                  className={`${fieldClass} h-32 font-mono`}
                  value={labBody}
                  onChange={(e) => setLabBody(e.target.value)}
                />
              </div>
              <div className="flex flex-wrap gap-3">
                <Button onClick={() => run("create-lab", () => apiPost("/labs", tryParseJson(labBody)))}>Create</Button>
                <Button variant="secondary" onClick={() => run("update-lab", () => apiPut(`/labs/${labId}`, tryParseJson(labBody)))}>
                  Update
                </Button>
                <Button variant="destructive" onClick={() => run("delete-lab", () => apiDelete(`/labs/${labId}`))}>
                  Delete
                </Button>
                <Button variant="outline" onClick={() => run("list-labs", () => apiGet("/labs"))}>
                  List
                </Button>
              </div>
            </div>
          </Panel>

          {/* Music */}
          <Panel title="Music / Playlists · CRUD / Query">
            <div className={sectionGap}>
              <div className="grid md:grid-cols-2 gap-3">
                <input
                  className={fieldClass}
                  placeholder="Music ID (update/delete)"
                  value={musicId}
                  onChange={(e) => setMusicId(e.target.value)}
                />
                <textarea
                  className={`${fieldClass} h-32 font-mono`}
                  value={musicBody}
                  onChange={(e) => setMusicBody(e.target.value)}
                />
              </div>
              <div className="flex flex-wrap gap-3 items-center">
                <Button onClick={() => run("create-music", () => apiPost("/admin/music", tryParseJson(musicBody)))}>
                  Create
                </Button>
                <Button
                  variant="secondary"
                  onClick={() => run("update-music", () => apiPut(`/admin/music/${musicId}`, tryParseJson(musicBody)))}
                >
                  Update
                </Button>
                <Button variant="destructive" onClick={() => run("delete-music", () => apiDelete(`/admin/music/${musicId}`))}>
                  Delete
                </Button>
                <Button variant="outline" onClick={() => run("list-music", () => apiGet("/music"))}>
                  List Music
                </Button>
                <input
                  className={`${fieldClass} w-36`}
                  placeholder="Playlist ID"
                  value={playlistId}
                  onChange={(e) => setPlaylistId(e.target.value)}
                />
                <Button variant="ghost" onClick={() => run("list-playlists", () => apiGet("/music/playlists"))}>
                  Playlists
                </Button>
                <Button variant="ghost" onClick={() => run("playlist-detail", () => apiGet(`/music/playlists/${playlistId}`))}>
                  Playlist Detail
                </Button>
              </div>
            </div>
          </Panel>

          <Card className="rounded-none border border-primary/30 bg-black/60">
            <CardContent className="p-4 space-y-2">
              <div className="flex items-center gap-2 text-primary font-bold uppercase tracking-widest">
                <Zap className="h-4 w-4" /> API Response
                {busy && (
                  <span className="flex items-center gap-1 text-xs text-muted-foreground">
                    <Loader2 className="h-4 w-4 animate-spin" />
                    {busy}
                  </span>
                )}
              </div>
              <pre className="text-xs whitespace-pre-wrap text-white/80 bg-black/40 p-3 border border-white/5 rounded-none max-h-96 overflow-auto">
                {message}
              </pre>
            </CardContent>
          </Card>
        </div>
      </section>
    </div>
  );
}
